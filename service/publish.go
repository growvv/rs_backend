package service

import (
	"github.com/growvv/rs_demo/dao"
	"github.com/growvv/rs_demo/model"
)

func Publish(name string, userId uint64) (model.Picture, bool) {

	pictureId, ok := dao.AddPicture(name, userId)
	if !ok {
		return model.Picture{}, false
	}
	picture := model.Picture{
		Id:   pictureId,
		Name: name,
	}
	return picture, ok
}

func PublishList(id uint64) []model.Picture {
	pictures := dao.GetPictureList(id)
	pictureInfos := make([]model.Picture, 0)
	for _, picture := range pictures {
		authorDb, _ := dao.GetUserById(picture.AuthorId)
		author := model.User{
			Id:   authorDb.Id,
			Name: authorDb.Username,
		}

		pictureInfos = append(pictureInfos, model.Picture{
			Id:     picture.Id,
			Name:   picture.Name,
			Author: author,
		})
	}
	return pictureInfos
}
