package common

import (
	"github.com/dgrijalva/jwt-go"
	"ginDemo/model"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {
	//token过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//获取jwt密钥
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "Arete",
			Subject:"user token",
		},
	}

	//使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(i interface{}, err error) {
		return jwtKey, nil
	})

	return token,claims,err
}