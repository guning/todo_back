package task

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	. "todo_back/handlers"
	"todo_back/models"
	"todo_back/pkg/constvar"
	"todo_back/pkg/errno"
)

func List(c *gin.Context) {
	log.Print("list tasks")

	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	r := ListRequest{
		TaskName: c.Query("taskName"),
		Limit: limit,
		Offset: offset,
	}


	if r.Limit == 0 {
		r.Limit = constvar.DefaultLimit
	}

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

	tasks, count, err := models.GetTaskList(u, r.TaskName, r.Offset, r.Limit)

	if err != nil {
		SendResponse(c, errno.ErrTaskList, err.Error())
		return
	}

	SendResponse(c, nil, ListResponse{
		TaskList: tasks,
		TotalCount: count,
	})

}
