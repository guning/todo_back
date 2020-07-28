package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

type Token struct {
	OpenId string `json:"openId"`
	UnionId string `json:"unionId"`
	ExpiredTime time.Time `json:"expiredTime"`
	jwt.StandardClaims
}

func (t Token) Valid() error {
	if !t.isExpired() {
		return errors.New("Token is expired.")
	}
	return nil
}

func (t *Token) isExpired() bool {
	if t.ExpiredTime.Unix() < time.Now().Unix() {
		return false
	}
	return true
}

func ParseValidateToken(tokenStr string) (Token, error) {
	var res Token
	token ,err := jwt.ParseWithClaims(tokenStr, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		parseBytes, err := ioutil.ReadFile(viper.GetString("key.public"))
		if err != nil {
			return []byte{}, err
		}
		parseKey, err := jwt.ParseRSAPublicKeyFromPEM(parseBytes)
		if err != nil {
			return []byte{}, err
		}
		return parseKey, nil
	})

	if err != nil {
		return res, err
	}

	if cliams, ok := token.Claims.(*Token); ok {
		log.Printf("token message is %s", cliams)
		return *cliams, nil
	} else {
		return res, err
	}
}


func SignToken(t Token) (string ,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, t)

	signBytes, err := ioutil.ReadFile(viper.GetString("key.private"))
	if err != nil {
		return "", err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return "", err
	}
	return token.SignedString(signKey)
}
