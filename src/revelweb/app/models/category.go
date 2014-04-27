package models

import (
	"time"
	// "github.com/revel/revel"
)

type Category struct {
	Id      int64
	Name    string			//总分类名称	//按类型:人物社会军事探索历史地理自然文化其它典藏奇艺旅游公开课
	Intro   string
	Cateid	int8			//分类id
	Created time.Time `xorm:"created"`
}



//重命名表
func (c *Category) TableName() string {
 	return "t_category"
 }


// func (category Category) Validate(v *revel.Validation) {
// 	v.Required(category.Name).Message("请输入名称")

// 	if category.HasName() {
// 		err := &revel.ValidationError{
// 			Message: "名称已存在",
// 			Key:     "category.Name",
// 		}
// 		v.Errors = append(v.Errors, err)
// 	}
// }

// func (c Category) HasName() bool {
// 	var category Category
// 	if c.Id > 0 {
// 		Engine.Where("name = ? AND id != ?", c.Name, c.Id).Get(&category)
// 	} else {
// 		Engine.Where("name = ?", c.Name).Get(&category)
// 	}

// 	return category.Id > 0
// }