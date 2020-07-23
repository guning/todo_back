package main

import (
	"github.com/spf13/pflag"
	. "todo_back/migration"
)

var (
	op = pflag.StringP("operation", "o", "rollUp", "sever config file path")
)

const (
	ROLLUP   string = "rollUp"
	ROLLDOWN string = "rollDown"
)


/*func main() {
	migrationList := []interface{}{
		DBInit{},
	}
	pflag.Parse()
	config.Init("")
	GetDBInstance()
	defer Close()
	DB.LogMode(true)
	MDBInit()
	if *op == ROLLUP {
		DoMigrate(migrationList)
	} else {
		DoRollDown(migrationList)
	}
}*/

func MDBInit() {
	if !DB.HasTable(&Migration{}) {
		DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Migration{})
		DB.Model(&Migration{}).ModifyColumn("version", "int not null default 0")
	}
}

func MigrationUp(name string) {
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

func MigrationDown() {
	record := Migration{}
	if err := DB.Model(&Migration{}).Last(&record).Error; err != nil {
		panic(err)
	}
	if err := DB.Model(&Migration{}).Where("Version = ?", record.Version).Delete(&Migration{}).Error; err != nil {
		panic(err)
	}
}

func DoMigrate(fs []interface{}) {
	for _, f := range fs {
		if v, ok := f.(Migrater); ok {
			defer func() {
				if r := recover(); r != nil {
					v.RollBack()
					MigrationDown()
				}
			}()
			v.RollUp()
			MigrationUp(v.GetName())
		}
	}
}

func DoRollDown(fs []interface{}) {
	record := Migration{}
	if err := DB.Model(&Migration{}).Last(&record).Error; err != nil {
		panic(err)
	}

	for _, f := range fs {
		if v, ok := f.(Migrater); ok && v.GetName() == record.Name {
			v.RollBack()
			MigrationDown()
		}
	}
}
