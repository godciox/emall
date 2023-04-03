package routers

import (
	"frontend/httphandler/user"
	"github.com/beego/beego/v2/server/web"
)

var HttpServer *web.HttpServer

func Route() {

	HttpServer.InsertFilter("/*", web.BeforeRouter, user.LoginFilter)
	// 用户登录
	HttpServer.Post("/user/login", user.LoginByPassword)

	// 用户登录
	HttpServer.Post("/user/login/captcha", user.LoginByCaptcha)

	// 用户获取验证码
	HttpServer.Post("/captcha", user.GetCaptcha)

	// 用户注册
	HttpServer.Post("/user/register", user.RegisterUser)

	// 获取用户信息
	HttpServer.Get("/user/getInfo", user.GetUserInfo)

	HttpServer.Router("/search/product", &user.ProductController{}, "post:SearchProduct")
}
