package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	IsLogin uint8
	UnionId string
	SessionKey string
	OpenId string
	Tasks []Task
}
