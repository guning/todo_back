package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
)

func Get(uri string, router *gin.Engine) []byte {
	req := httptest.NewRequest("GET", uri, nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body
}


func PostJson(uri string, param map[string]interface{}, router *gin.Engine, method string) []byte {
	jsonByte, _ := json.Marshal(param)

	req := httptest.NewRequest(method, uri, bytes.NewReader(jsonByte))

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return body
}
