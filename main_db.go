package main

import (
	"github.com/spf13/pflag"
	"todo_back/config"
	. "todo_back/migration"
)

var (
	op = pflag.StringP("operation", "o", "rollUp", "sever config file path")
)

const (
	ROLLUP   string = "rollUp"
	ROLLDOWN string = "rollDown"
)


func main() {
	migrationList := []interface{}{
		DBInit{},
	}
	pflag.Parse()
	config.Init("")
	GetDBInstance()
	defer Close()
	DB.LogMode(true)
	dbInit()
	if *op == ROLLUP {
		doMigrate(migrationList)
	} else {
		doRollUp(migrationList)
	}
}

func dbInit() {
	if !DB.HasTable(&Migration{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Migration{})
		DB.Model(&Migration{}).ModifyColumn("version", "int not null default 0")
	}
}

func migrationUp(name string) {
	migrate := Migration{
		Name: name,
	}
	lastRecord := Migration{}
	if err := DB.Model(&Migration{}).Last(&lastRecord).Error; err != nil {
		panic(err)
	}

	if lastRecord.Name == name {
		return
	}

	migrate.Version = lastRecord.Version + 1
	if err := DB.Model(&Migration{}).Create(&migrate).Error; err != nil {
		panic(err)
	}
}

func migrationDown() {
	record := Migration{}
	if err := DB.Model(&Migration{}).Last(&record).Error; err != nil {
		panic(err)
	}
	if err := DB.Model(&Migration{}).Where("Version = ?", record.Version).Delete(&Migration{}).Error; err != nil {
		panic(err)
	}
}

func doMigrate(fs []interface{}) {
	for _, f := range fs {
		if v, ok := f.(Migrater); ok {
			defer func() {
				if r := recover(); r != nil {
					v.RollBack()
					migrationDown()
				}
			}()
			v.RollUp()
			migrationUp(v.GetName())
		}
	}
}

func doRollUp(fs []interface{}) {
	record := Migration{}
	if err := DB.Model(&Migration{}).Last(&record).Error; err != nil {
		panic(err)
	}

	for _, f := range fs {
		if v, ok := f.(Migrater); ok && v.GetName() == record.Name {
			v.RollBack()
			migrationDown()
		}
	}
}
