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

//用户列表
func (c Admin) UserList() revel.Result {
	var users []*models.User
	engine.Find(&users)
	return c.Render(users)
}

//删除用户
func(c Admin) UserDelete(id int64) revel.Result{
	aff,_:= engine.Id(id).Delete(&models.User{})
	if aff > 0{
		//TODO ajax aff
	}
	//TODO ajax
	return c.Redirect(routes.Admin.UserList())
}

//激活用户
func(c Admin) UserActivate(id int64) revel.Result{
	aff,_:=engine.Id(id).Cols("status").Update(&models.User{
		Status:models.USER_STATUS_ACTIVATED,
	})

	if aff > 0{
		//TODO aff
	}
	return c.Redirect(routes.Admin.UserList())
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


//文章列表
func(c Admin) ArticleList() revel.Result{
	var articles []*models.Article
	engine.Find(&articles)
	return c.Render(articles)
}

//添加文章页
func(c Admin) ArticleNew() revel.Result{
	title := "新建文章"
	//TODO category
	return c.Render(title)
}

//添加
func (c Admin) ArticlePost(article models.Article) revel.Result{
	article.Hits = 0
	aff,_:=engine.Insert(&article)
	if aff > 0{
		//TODO aff
	}
	return c.Redirect(routes.Admin.ArticleList())
}


//删除
func (c Admin) ArticleDelete(id int64) revel.Result{
	aff,_:=engine.Id(id).Delete(&models.Article{})
	if aff > 0{
		//TODO aff
	}
	return c.Redirect(routes.Admin.ArticleList())
}

//修改文章
func (c Admin) ArticleEdit(id int64) revel.Result{
	title:="编辑文章"	//和添加文章用一个模板
	var article models.Article
	has,_:=engine.Id(id).Get(&article)
	if !has{
		return c.NotFound("文章不存在")
	}

	c.vars(Vars{
		"title":title,
		"article":article,
	})
	return c.RenderTemplate("admin/articlenew.html")
}

//修改文章保存
func (c Admin) ArticleEditPost(id int64,article models.Article) revel.Result{
	article.Id = id
	aff,_:=engine.Id(id).Update(&article)
	if aff > 0{
		//TODO aff
	}
	return c.Redirect(routes.Admin.ArticleList())	
}