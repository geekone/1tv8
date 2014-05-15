package models
//文章模型

import (
	"time"
)


type Article struct {
	Id int64		//文章编号
	Title string						//标题
	Intro	string `xorm:"text"`		//介绍
	Content	string `xorm:"text"`		//内容
	Hits int64							//点击量
	Created time.Time `xorm:"created"`	//建立时间
	Updated time.Time `xorm:"updated"`	//更新日期
}

//重命名表名
func (m *Article) TableName() string {
	return "t_article"
}