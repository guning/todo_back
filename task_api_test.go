package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
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
	r.POST("/createTask", task.Create)
}

func Get(uri string, router *gin.Engine) []byte {
	req := httptest.NewRequest("GET", uri, nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body
}


func PostJson(uri string, param map[string]interface{}, router *gin.Engine) []byte {
	jsonByte, _ := json.Marshal(param)

	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func TestTaskCreate(t *testing.T) {
	uri := "/createTask"

	param := make(map[string]interface{})

	param["unionId"] = "unionId"
	param["taskName"] = "taskName"
	param["deadline"] = time.Now()
	param["detail"] = "detail"

	body := PostJson(uri, param, r)

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