package routers

import (
	"frontend/httphandler"
	"github.com/beego/beego/v2/server/web"
)

var HttpServer *web.HttpServer

func Route() {

	HttpServer.InsertFilter("/*", web.BeforeRouter, httphandler.LoginFilter)
	// 用户登录
	HttpServer.Post("/user/login", httphandler.LoginByPassword)

	// 用户注册
	HttpServer.Post("/user/register", httphandler.RegisterUser)

	// 获取用户信息
	HttpServer.Get("/user/getInfo", httphandler.GetUserInfo)

}
