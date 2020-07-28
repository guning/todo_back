package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
	"todo_back/clients/auth"
	. "todo_back/handlers"
	"todo_back/models"
	"todo_back/pkg/errno"
	"todo_back/pkg/token"
)

func Auth(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		SendResponse(c, errno.ErrCodeInvalid, nil)
	}

	r := auth.Request{
		AppId: viper.GetString("appid"),
		Secret: viper.GetString("secret"),
		JsCode: code,
	}

	resp, err := auth.Code2Session(r)
	if err != nil {
		SendResponse(c, errno.ErrGetUser, nil)
		return
	}

	m, _ := time.ParseDuration("1m")
	t, err := token.SignToken(token.Token{
		OpenId: resp.OpenId,
		UnionId: resp.UnionId,
		ExpiredTime: time.Now().Add(30 * m),
	})

	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	u := models.User{
		OpenId: resp.OpenId,
		UnionId: resp.UnionId,
	}

	if err := u.Create(); err != nil {
		log.Print("user create err:", err)
		SendResponse(c, errno.ErrUserCreate, nil)
		return
	}

	SendResponse(c, nil, t)
}
