package models

import (
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

func (t *Task) Delete() error {
	return DB.Self.Delete(&t).Error
}

func GetTaskList(u User, taskName string, offset, limit int) ([]*Task, uint64, error) {
	if limit == 0 {
		limit = 10
	}

	tasks := make([]*Task, 0)
	var count uint64

	tmp := DB.Self.Model(&Task{})
	tmp = tmp.Where("userId = ?", u.ID)
	if taskName != "" {
		tmp = tmp.Where("taskName like ?", "%" + taskName + "%")
	}


	if err := tmp.Count(&count).Error; err != nil {
		return tasks, count, err
	}

	if err := tmp.Offset(offset).Limit(limit).Find(&tasks).Error; err != nil {
		return tasks, count, err
	}
	return tasks, count, nil
}
