package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	UserId uint `gorm:"index"`
	TaskName string
	Detail string
	Deadline time.Time
}
