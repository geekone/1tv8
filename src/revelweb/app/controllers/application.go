package controllers

/**
 * 继承的主类
 */

import (
	"github.com/revel/revel"
	"github.com/go-xorm/xorm"
)

var (
	engine      *xorm.Engine
	
)


type Vars map[string]interface{}

type Application struct {
	*revel.Controller
}

//多个变量到页面里
func (c *Application) vars(vars Vars) {
	for k, v := range vars {
		c.RenderArgs[k] = v
	}
}