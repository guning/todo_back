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

func Delete(c *gin.Context) {
	log.Print("delete task")
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

	log.Printf("Task id is %d", id)

	tmp, ok := c.Get("user")
	if !ok {
		SendResponse(c, errno.ErrUserNotFound, "cannot get user")
		return
	}
	u, ok := tmp.(models.User)

	if !ok {
		SendResponse(c, errno.ErrUserNotFound, "invalid user")
		return
	}

	t := models.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
		UserId: u.ID,
	}

	if err := t.Delete(); err != nil {
		SendResponse(c, errno.ErrTaskDelete, nil)
		return
	}

	SendResponse(c, nil, CreateResponse{
		TaskId: t.ID,
	})
}
