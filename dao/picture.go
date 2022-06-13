package dao

import (
	"github.com/growvv/rs_demo/config"
	"github.com/growvv/rs_demo/model"
)

func AddPicture(name string, userId uint64) (uint64, bool) {
	postPicture := model.PictureDB{
		Name:     name,
		AuthorId: userId,
	}
	res := config.Db.Create(&postPicture)
	if res.Error != nil {
		return 0, false
	}
	return postPicture.Id, true
}

func GetPictureList(userId uint64) []model.PictureDB {
	var pictures []model.PictureDB
	config.Db.Model(&model.PictureDB{}).Where("author_id = ?", userId).Find(&pictures)
	return pictures
}
