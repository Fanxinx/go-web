package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 24 * 7

var mySecret = []byte("nibuzhidaodemimi")

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

func GenToken(userId int64,username string) (string,error){
	c := MyClaims{
		UserId: userId,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),//过期时间
			Issuer:	   "webapp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*MyClaims,error){

	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString,mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret,nil
	})

	if err != nil{
		return nil,err
	}

	if token.Valid{
		return mc,nil
	}

	return nil,errors.New("invalid token")
}