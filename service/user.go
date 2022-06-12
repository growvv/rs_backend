package service

import (
	"github.com/growvv/rs_demo/dao"
	"github.com/growvv/rs_demo/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte = []byte("xxxxxx")

type Claims struct {
	Id uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(Id uint64) (string, error) {
	claims := Claims{Id, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 2*60*60,
		Issuer:    "LFR",
	},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func Register(username, password string) (uint64, bool) {
	count := dao.ExistUser(username)
	if count > 0 { // 如果用户名已存在
		return 0, false
	}
	id, ok := dao.AddUser(username, password)
	if ok {
		return id, true
	} else {
		return 0, false
	}
}

func Login(name string, password string) (model.User, bool) {
	userDb, ok := dao.GetUser(name, password)
	user := model.User{
		Id:   userDb.Id,
		Name: userDb.Username,
	}
	return user, ok
}
