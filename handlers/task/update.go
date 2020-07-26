package task

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	. "todo_back/handlers"
	"todo_back/models"
	"todo_back/pkg/errno"
)

func Update(c *gin.Context) {
	log.Print("update task")
	var r CreateRequest

	if err := c.BindJSON(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	strId := c.Param("id")
	id ,err := strconv.Atoi(strId)
	if err != nil {
		SendResponse(c, errno.ErrTaskNotFound, nil)
		return
	}

	log.Printf("Task id is %d, unionId is %s", id, r.UnionId)

	u, err := models.FindByUnionId(r.UnionId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	t := models.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
		TaskName: r.TaskName,
		Detail: r.Detail,
		UserId: u.ID,
		Deadline: r.Deadline,
	}

	if err := t.Update(); err != nil {
		SendResponse(c, errno.ErrTaskUpdate, nil)
		return
	}

	SendResponse(c, nil, CreateResponse{
		TaskId: t.ID,
	})
}
