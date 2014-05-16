package controllers

import(
	"github.com/revel/revel"
	"revelweb/app/routes"
	"revelweb/app/models"
)

type User struct{
	Application
}

//注册
func (c User) Signup() revel.Result {
	return c.Render()
}

//注册保存
func (c User) SignupPost(user models.User) revel.Result {
	user.Validate(c.Validation)			//用户的验证方法来调用revel controller的验证对象,然后返回结果
	
	//是否有错误
	if c.Validation.HasErrors(){
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.User.Signup())		//如果出错返回注册页
	}

	//通过验证插入数据库
	aff,_ := engine.Insert(&models.User{
		Name:	user.Name,				//用户名
		Email:	user.Email,				//邮件
		Type:	MEMBER_GROUP,			//权限
		Avatar:	models.DefaultAvatar,	//头象
		ValidateCode: "abc",			//验证码
		Salt:	"salt",					
		HashedPassword: user.Password,	//TODO 没加密码
	})

	if aff == 0{
		//TODO 注册失败后
	}


	return c.Redirect(routes.User.Signin())			//成功去登录页
}

//用户登录页
func (c User) Signin() revel.Result{
	return c.Render()
}