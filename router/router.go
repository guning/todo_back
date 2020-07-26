package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_back/handlers/task"
	. "todo_back/router/middlewares"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(NoCache)
	g.Use(Options)
	g.Use(Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "path not found"})
	})

	u := g.Group("/user")

	{
		u.POST("/login")
		u.POST("/login/quit")
	}

	t := g.Group("/task")

	{
		t.POST("", task.Create) //new
		t.DELETE("/:id") //delete
		t.PUT("/:id", task.Update) //update
		t.GET("") //list
	}
	return g
}
