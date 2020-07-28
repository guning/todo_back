package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_back/config"
	"todo_back/handlers"
	"todo_back/handlers/task"
	"todo_back/models"
)

var r *gin.Engine

func init() {
	config.Init("")
	models.DB.Init()
	models.DB.Self.LogMode(true)
	r = gin.Default()
	r.Use(tokenMiddleware)
	r.POST("/createTask", task.Create)
	r.PUT("/updateTask/:id", task.Update)
	r.DELETE("/deleteTask/:id", task.Delete)
	r.GET("", task.List)
}

func tokenMiddleware(c *gin.Context) {
	c.Set("user", models.User{
		UnionId: "unionId",
		Model: gorm.Model{
			ID: 1,
		},
	})
	c.Next()
}



func TestTaskCreate(t *testing.T) {
	uri := "/createTask"

	param := make(map[string]interface{})

	param["taskName"] = "taskName"
	param["deadline"] = time.Now()
	param["detail"] = "detail"
	createUpdateDelete(t, uri, param, "POST")

}

func TestTaskUpdate(t *testing.T) {
	uri := "/updateTask/9"

	param := make(map[string]interface{})

	param["taskName"] = "taskName27"
	param["deadline"] = time.Now()
	param["detail"] = "detail1"
	createUpdateDelete(t, uri, param, "PUT")
}

func TestTaskDelete(t *testing.T) {
	uri := "/deleteTask/6"

	param := make(map[string]interface{})

	createUpdateDelete(t, uri, param, "DELETE")
}

func TestTaskList(t *testing.T) {
	//uri := "/?limit=10&unionId='+or+1=1"
	uri := "/?limit=10&taskName=taskName2"


	body := Get(uri, r)
	fmt.Printf("response: %v \n", string(body))

	res := handlers.Response{}
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println("json unmarshal failed")
		panic(err)
	}

	assert.Equal(t, 0 , res.Code)
	assert.Equal(t, "OK", res.Message)
	data := res.Data
	if m, ok := data.(map[string]interface{}); ok {
		if count, ok := m["totalCount"]; ok {
			fmt.Println(count)
			return
		}
	}
	panic("parse response failed")
}


func createUpdateDelete(t *testing.T, uri string, param map[string]interface{}, method string) {
	body := PostJson(uri, param, r, method)

	fmt.Printf("response: %v \n", string(body))

	res := handlers.Response{}
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println("json unmarshal failed")
		panic(err)
	}

	assert.Equal(t, 0 , res.Code)
	assert.Equal(t, "OK", res.Message)
	data := res.Data
	if m, ok := data.(map[string]interface{}); ok {
		if taskId, ok := m["taskId"]; ok {
			assert.IsType(t, float64(1), taskId)
			return
		}
	}
	panic("parse response failed")
}

