package controllers

import (
	"github.com/revel/revel"
	// "fmt"
	"revelweb/app/models"
)

type App struct {
	Application
}

func (c App) Index() revel.Result {
	var movies []*models.Movie
	var cates []*models.Category
	rows, _ := engine.Count(&models.Movie{})
	engine.Find(&movies)
	engine.Find(&cates)

	return c.Render(rows,movies,cates)
}


