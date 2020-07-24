package task

import (
	"time"
	"todo_back/models"
)

type CreateRequest struct {
	UnionId string `json:"unionId"`
	TaskName string `json:"taskName"`
	Detail string `json:"detail"`
	Deadline time.Time `json:"deadline"`
}

type CreateResponse struct {
	TaskId uint `json:"taskId"`
}

type ListRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	TaskList   []*models.Task `json:"taskList"`
}
