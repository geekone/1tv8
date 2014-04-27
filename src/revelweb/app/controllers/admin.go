package controllers

import (
	"github.com/revel/revel"
	"revelweb/app/models"
	"revelweb/app/routes"
	// "fmt"
)

type Admin struct {
	Application
}

func (c Admin) Index() revel.Result{
	return c.Render()
}


//分类列表
func (c Admin) CategoryList() revel.Result{
	var categories []*models.Category
	engine.Find(&categories)
	return c.Render(categories)
}

//添加分类
func (c Admin) CategoryNew() revel.Result{
	title:="新建分类"
	return c.Render(title)
}

//新增保存
func (c Admin) CategoryPost(category models.Category) revel.Result{
	//TODO 验证
	

	aff,_:=engine.Insert(&category)
	if aff > 0{
		//TODO 成功之后
	}

	return c.Redirect(routes.Admin.CategoryList())
}

//删除分类
func (c Admin) CategoryDelete(id int64) revel.Result{
	aff,_:=engine.Id(id).Delete(&models.Category{})
	if aff > 0{
		//TODO 删除成功之后
	}
	return c.Redirect(routes.Admin.CategoryList())
}

//跳到修改页面，使用新建一样的页面
func (c Admin) CategoryEdit(id int64) revel.Result{
	title := "编辑分类"
	var category models.Category
	has, _ := engine.Id(id).Get(&category)
	if !has {
		return c.NotFound("分类不存在")
	}

	c.vars(Vars{
		"title":    title,
		"category": category,
	})

	return c.RenderTemplate("admin/categorynew.html")
}

//修改保存分类
func (c Admin) CategoryEditPost(id int64,category models.Category) revel.Result{
			category.Id = id
			//TODO 验证
			aff,_:=engine.Id(id).Update(&category)
			if aff > 0{
				//TODO 修改成功之后
			}
			return c.Redirect(routes.Admin.CategoryList())
}



//视频列表
func (c Admin) MovieList() revel.Result{
	var movies []*models.Movie
	engine.Find(&movies)
	return c.Render(movies)
}

//添加视频页面
func (c Admin) MovieNew() revel.Result{
	title := "新建视频"
	
	var categories []*models.Category
	engine.Find(&categories)		//TODO 这个可以放在cache或者application vars
	return c.Render(title,categories)
}

//添加视频保存
func (c Admin) MoviePost(movie models.Movie,category int64) revel.Result {
	movie.Category = models.Category{Id:category}
	aff,_:=engine.Insert(&movie)
	if aff > 0 {
		//TODO 成功之后
	}
	return c.Redirect(routes.Admin.MovieList())
}

//删除视频
func (c Admin) MovieDelete(id int64) revel.Result {
	aff,_:= engine.Id(id).Delete(&models.Movie{})
	if aff > 0{
		//TODO 删除成功之后
	}
	return c.Redirect(routes.Admin.MovieList())
}

//跳转到编辑视频
func (c Admin) MovieEdit(id int64) revel.Result{
	title:="编辑视频"
	var movie models.Movie
	has,_:=engine.Id(id).Get(&movie)
	if !has{
		return c.NotFound("视频不存在")
	}

	c.vars(Vars{
		"title":title,
		"movie":movie,
	})

	return c.RenderTemplate("admin/movienew.html")
}

//修改之后保存视频
func (c Admin) MovieEditPost(id int64,movie models.Movie) revel.Result{
	movie.Id = id
	aff,_:=engine.Id(id).Update(&movie)
	if aff > 0{
		//TODO 修改成功之后
	}
	return c.Redirect(routes.Admin.MovieList())
}



//测试preview
func (c Admin) MoviePreview(id int64) revel.Result {


	var movie models.Movie
	has, _ := engine.Id(id).Get(&movie)
	if !has {
		return c.NotFound("分类不存在")
	}

	return c.Render(movie)

}