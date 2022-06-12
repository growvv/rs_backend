package dao

import (
	"github.com/growvv/rs_demo/config"
	"github.com/growvv/rs_demo/model"
)

func ExistUser(username string) int64 {
	var count int64
	config.Db.Model(&model.UserDB{}).Where("username = ?", username).Count(&count)
	return count
}

func AddUser(name, password string) (uint64, bool) {
	user := model.UserDB{
		Username: name,
		Password: password,
	}
	row := config.Db.Create(&user).RowsAffected
	if row == 0 {
		return 0, false
	} else {
		return user.Id, true
	}
}

func GetUser(username, password string) (model.UserDB, bool) {
	var user model.UserDB
	config.Db.Where("username = ? and password = ?", username, password).First(&user) // 查询用户
	if user.Id == 0 {
		return model.UserDB{}, false
	} else {
		return user, true
	}
}
