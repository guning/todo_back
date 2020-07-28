package auth

import (
	"fmt"
	"github.com/gin-gonic/gin/internal/json"
	"github.com/google/go-querystring/query"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type Request struct {
	AppId string
	Secret string
	JsCode string
	GrantType string
}

type Response struct {
	OpenId string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId string `json:"unionid"`
	Errcode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func Code2Session(r Request) (Response, error){
	var res Response
	uri := "/sns/jscode2session"
	r.GrantType = "authorization_code"

	v, _ := query.Values(r)
	url := fmt.Sprintf("%s%s?%s", viper.GetString("miniProgram.host"), uri, v)
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, nil
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return res, err
	}

	return res, nil

}