package model

type UserDB struct {
	Id       uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

type User struct {
	Id   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
