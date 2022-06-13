package controller

import (
	"github.com/growvv/rs_demo/response"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/growvv/rs_demo/model"
	"github.com/growvv/rs_demo/service"
)

var usersLoginInfo map[string]model.User = make(map[string]model.User)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	log.Println("register: ", "username: ", username, " password: ", password)

	if username == "" || password == "" {
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Username or Password is empty"},
		})
		return
	}

	// token := username + password
	uid, ok := service.Register(username, password)
	if !ok {
		log.Println("Register fail, username already exist")
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Register fail"},
		})
		return
	}

	token, err := service.GenerateToken(uid)
	if err != nil {
		log.Println("Generate token fail")
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}

	newUser := model.User{
		Id:   uid,
		Name: username,
	}
	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, response.UserLoginResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   uid,
		Token:    token,
	})

}

// Login 用户登录接口
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	log.Print("login: ", "username: ", username, " password: ", password)

	if username == "" || password == "" {
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Username or Password is empty"},
		})
		return
	}

	user, ok := service.Login(username, password)
	if !ok {
		log.Println("Login fail")
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Login fail! Check your username and password"},
		})
		return
	}
	// token := username + password
	token, err := service.GenerateToken(user.Id)
	if err != nil {
		log.Println("Generate token fail", err)
		c.JSON(http.StatusOK, response.UserLoginResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Generate token fail"},
		})
		return
	}
	usersLoginInfo[token] = user
	c.JSON(http.StatusOK, response.UserLoginResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "Login Successfully"},
		UserId:   user.Id,
		Token:    token,
	})
}

// UserInfo 获取当前登录用户信息
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "No Token!"})
		return
	} else {
		claims, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "Parse Token fail!"})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "Token Timeout! Please login again"})
			return
		}
	}

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, response.UserResponse{
			Response: response.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, response.UserResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "Please Login again"},
		})
	}
}
