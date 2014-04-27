package controllers

/**
 * 初始化的类，和beego不一样，它是放在base.go里面
 */

import (
 	"github.com/revel/revel"
 	"revelweb/app/models"
)


func init(){
	revel.OnAppStart(Init)
	
}

//OnAppStart调用的初始程序
func Init(){
	engine = models.Engine
}