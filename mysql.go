package main

import (
	"fmt"

	"github.com/growvv/rs_demo/config"
	"github.com/growvv/rs_demo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initMysql() {
	var dsn = fmt.Sprintf("%s:%s@%s", config.MysqlUsername, config.MysqlPassword, config.MysqlUrl)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Errorf("gorm open error: %s\n", err))
	}

	// Todo 加其他model
	if err := db.AutoMigrate(&model.UserDB{}); err != nil {
		panic(fmt.Errorf("db automigrate err: %s\n", err))
	}

	config.Db = db
}
