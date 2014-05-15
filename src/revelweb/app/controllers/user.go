package controllers

import(
	"github.com/revel/revel"
	"revelweb/app/routes"
)

type User struct{
	Application
}

//注册
func (c User) Signup() revel.Result {
	return c.Render()
}

//注册保存
func (c User) SignupPost() revel.Result {
	return c.Redirect(routes.User.Signup())
}