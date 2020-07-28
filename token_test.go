package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"todo_back/config"
	"todo_back/pkg/token"
)

func TestToken(t *testing.T) {
	config.Init("")
	m, _ := time.ParseDuration("1m")
	tk := token.Token{
		OpenId:      "openId",
		UnionId:     "unionId",
		ExpiredTime: time.Now().Add(30 * m),
		StandardClaims: jwt.StandardClaims{
			Issuer: "test",
		},
	}

	fmt.Println("sign")
	tokenString, err := token.SignToken(tk)

	assert.NoError(t, err)

	fmt.Println(tokenString)

	fmt.Println("parse")
	tmp, err := token.ParseValidateToken(tokenString)

	assert.NoError(t, err)
	fmt.Println(tmp)
}
