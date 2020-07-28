package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_back/handlers"
	"todo_back/models"
	"todo_back/pkg/errno"
	t "todo_back/pkg/token"
)


func TokenH(c *gin.Context) {
	if c.Request.RequestURI == "/auth" {
		c.Next()
	}

	token := c.Request.Header.Get("AuthToken")
	if token == "" {
		c.AbortWithStatus(401)
	} else {
		if t, err := t.ParseValidateToken(token); err != nil {
			log.Print("token invalid", err)
			c.AbortWithStatus(403)
		} else {
			u, err := models.FindByUnionId(t.UnionId)
			if err != nil {
				handlers.SendResponse(c, errno.ErrUserNotFound, nil)
				return
			}
			c.Set("user", u)
		}
	}
	c.Next()
}


