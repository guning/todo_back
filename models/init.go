package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self:   GetSelfDB(),
	}
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}


func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.name"))
}


func openDB(username, password, host, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		host,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatal("DB connected failed: "+config, err)
	}

	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
}

func (db *Database) Close() {
	DB.Self.Close()
}