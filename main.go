package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"todo_back/config"
	"todo_back/models"
	"todo_back/router"
	"todo_back/router/middlewares"
)

var (
	cfg = pflag.StringP("config", "c", "", "sever config file path")
)

func main() {
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	router.Load(g, []gin.HandlerFunc{
		middlewares.TokenH,
	}...)

	models.DB.Init()
	defer models.DB.Close()

	log.Printf("start app, listening on %s", viper.GetString("port"))
	log.Printf(http.ListenAndServe(":" + viper.GetString("port"), g).Error())
}
