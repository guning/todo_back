package migration

import (
	"github.com/jinzhu/gorm"
	"todo_back/models"
)

type Migrater interface {
	RollUp()
	RollBack()
	GetName() string
}

var DB *gorm.DB

func GetDBInstance() *gorm.DB {
	if DB == nil {
		DB = models.GetSelfDB()
	}
	return DB
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}

type Migration struct {
	ID int `gorm:"primary_key"`
	Name string
	Version int `gorm:"DEFAULT 0"`
}