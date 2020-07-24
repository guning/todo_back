package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	IsLogin uint8
	UnionId string
	SessionKey string
	OpenId string
	Tasks []Task
}

func FindByUnionId(unionId string) (User, error) {
	u := User{}
	if err := DB.Self.Where("unionId = ?", unionId).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
