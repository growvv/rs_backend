package model

type PictureDB struct {
	Id         uint64 `gorm:"column:id;autoIncrement;primaryKey"`
	Name       string `gorm:"column:name"`
	AuthorId   uint64 `gorm:"column:author_id"`
	CreateTime int64  `gorm:"column:create_time"`
}

type Picture struct {
	Id     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author User   `json:"author,omitempty"`
}
