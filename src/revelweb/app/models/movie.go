package models

import (
	"time"
)

type Movie struct {
	Id int64
	Title string
	Img   string
	Intro	string `xorm:"text"`
	Content string `xorm:"text"`
	Category Category `xorm:"category_id bigint"`
	Hits	int
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

//重命名表
func (m *Movie) TableName() string{
	return "t_movie"
}