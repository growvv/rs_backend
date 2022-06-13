package controller

import (
	"fmt"
	"github.com/growvv/rs_demo/response"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/growvv/rs_demo/service"
)

// Publish 上传图片
func Publish(c *gin.Context) {

	token, ok := c.GetPostForm("token")
	token2 := c.PostForm("token")
	log.Println(token2)
	if !ok {
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  "Public Request Token Fail",
		})
		return
	}

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

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, response.Response{StatusCode: 1, StatusMsg: "Please login again"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		log.Println("data error!\n")
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user := usersLoginInfo[token]
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%d_%s", user.Id, time.Now().Nanosecond(), filename)
	savePath := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		log.Println("save error!")
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	} else {
		picture, ok := service.Publish(finalName, user.Id)
		if ok {
			c.JSON(http.StatusOK, response.PictureResponse{
				Response: response.Response{StatusCode: 0, StatusMsg: filename + " uploaded successfully"},
				Picture:  picture,
			})
		} else {
			c.JSON(http.StatusOK, response.PictureResponse{
				Response: response.Response{StatusCode: 1, StatusMsg: "Upload Fail"},
			})
		}
	}
}

// PublishList 获取视频列表
func PublishList(c *gin.Context) {
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

	userPictureList := service.PublishList(usersLoginInfo[token].Id)
	c.JSON(http.StatusOK, response.PictureListResponse{
		Response: response.Response{
			StatusCode: 0,
		},
		PictureList: userPictureList,
	})
}
