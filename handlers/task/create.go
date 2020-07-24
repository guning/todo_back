package task

import (
	"github.com/gin-gonic/gin"
	"log"
	. "todo_back/handlers"
	"todo_back/models"
	"todo_back/pkg/errno"
)

func Create(c *gin.Context) {
	log.Printf("create task")
	var r CreateRequest
	if err := c.BindJSON(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	log.Printf("user id is %s, task msg is %s", r.UnionId, r.TaskName)

	u, err := models.FindByUnionId(r.UnionId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	t := models.Task{
		UserId: u.ID,
		TaskName: r.TaskName,
		Deadline: r.Deadline,
		Detail: r.Detail,
	}

	if err := t.Create(); err != nil {
		log.Print(err)
		SendResponse(c, errno.ErrTaskCreate, nil)
		return
	}

	SendResponse(c, nil, CreateResponse{
		TaskId: t.ID,
	})
}
