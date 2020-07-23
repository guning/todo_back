package main

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
	"todo_back/config"
	. "todo_back/migration"
)

type MT struct{}

func (mt MT) RollUp() {}

func (mt MT) RollBack() {}

func (mt MT) GetName() string{return "mt test"}

func TestMigration(t *testing.T) {
	config.Init("")
	GetDBInstance()
	defer Close()

	//DB.LogMode(true)
	if !DB.HasTable("migrations") {
		panic(errors.New("migrations not found"))
	}
	oldMigration := Migration{}
	if err := DB.Model(&Migration{}).Last(&oldMigration).Error; err != nil {
		panic(err)
	}
	migrationList := []interface{}{
		MT{},
	}
	DoMigrate(migrationList)
	newMigration := Migration{}
	if err := DB.Model(&Migration{}).Last(&newMigration).Error; err != nil {
		panic(err)
	}
	assert.Equal(t, newMigration.Version, oldMigration.Version + 1, "Version err 1")

	DoRollDown(migrationList)
	newMigration = Migration{}

	if err := DB.Model(&Migration{}).Last(&newMigration).Error; err != nil {
		panic(err)
	}
	assert.Equal(t, newMigration.Version, oldMigration.Version, "Version err 2")
}