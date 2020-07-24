package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	UserId uint `gorm:"index"`
	TaskName string `gorm:""`
	Detail string
	Deadline time.Time
}

func (t *Task) Create() error {
	return DB.Self.Create(&t).Error
}

func (t *Task) Update() error {
	return DB.Self.Save(&t).Error
}

func DeleteById(id uint) error {
	task := Task{
		Model:    gorm.Model{ID: id},
	}
	return DB.Self.Delete(&task).Error
}

func GetTaskList(taskName string, offset, limit int) ([]*Task, uint64, error) {
	if limit == 0 {
		limit = 10
	}

	tasks := make([]*Task, 0)
	var count uint64

	where := fmt.Sprintf("taskName like '%%%s%%'", taskName)

	if err := DB.Self.Model(&Task{}).Where(where).Count(&count).Error; err != nil {
		return tasks, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return tasks, count, err
	}
	return tasks, count, nil
}
