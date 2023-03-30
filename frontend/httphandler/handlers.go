package httphandler

import (
	"context"
	"encoding/json"
	"fmt"
	proto "frontend/proto"
	"frontend/service"
	"frontend/utils/token"
	context2 "github.com/beego/beego/v2/server/web/context"
	"net/http"
)

type UserLoginParams struct {
	Mobile   string `form:"mobile"`
	Password string `form:"password"`
	UserName string `form:"username"`
}

func LoginByPassword(ctx *context2.Context) {
	var us UserLoginParams
	if err := ctx.BindForm(&us); err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		ctx.WriteString("错误请求")
		return
	}
	var user *proto.UserResponse
	var err error
	if us.UserName != "" {
		user, err = service.Svc.UserService.CheckPasswordToUser(context.Background(), &proto.User{
			Username: us.UserName,
			Password: us.Password,
		})
	} else {
		user, err = service.Svc.UserService.LoginByMobile(context.Background(), &proto.User{
			Mobile:   us.Mobile,
			Password: us.Password,
		})
	}
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		ctx.WriteString(fmt.Errorf("出错信息：%s", err).Error())
		return
	}
	if user.Status == "500" {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	var tokenString string
	if us.Mobile != "" {
		tokenString = token.CreateToken(us.Mobile)
	} else {
		tokenString = token.CreateToken(us.UserName)
	}

	ctx.ResponseWriter.Header().Add("Authorization", tokenString)
}

type RegisterUserParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Name     string `form:"name"`
	Email    string `form:"email"`
	Mobile   string `form:"mobile"`
	Gender   int32  `form:"gender"`
}

func RegisterUser(ctx *context2.Context) {
	var in RegisterUserParams
	if err := ctx.BindForm(&in); err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		ctx.WriteString("错误请求,注册格式不对" + err.Error())
		return
	}

	rsp, _ := service.Svc.UserService.RegisterUser(context.Background(), &proto.User{
		Username: in.Username,
		Password: in.Password,
		Name:     in.Name,
		Email:    in.Email,
		Mobile:   in.Mobile,
		Gender:   in.Gender,
	})
	if rsp.Status == "500" {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		ctx.WriteString(rsp.String())
		return
	}
	ctx.ResponseWriter.WriteHeader(http.StatusOK)
	return
}

type GetUserInfoParams struct {
	UserName string `form:"username"`
	Mobile   string `form:"mobile"`
}

func GetUserInfo(ctx *context2.Context) {
	var in GetUserInfoParams
	tokenS := ctx.Request.Header["Authorization"]
	mobile, _ := token.CheckToken(tokenS[0])
	in.Mobile = mobile
	var us = new(proto.User)
	us.Mobile = in.Mobile
	user, _ := service.Svc.UserService.GetUserByPhone(context.Background(), us)
	if user.Response.Status == "500" {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		ctx.WriteString(user.Response.String())
		return
	}
	byteArr, err := json.Marshal(user)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.WriteString(string(byteArr))
}

func LoginFilter(ctx *context2.Context) {
	tokenS := ctx.Request.Header["Authorization"]

	if ctx.Request.URL.Path == "/user/register" || ctx.Request.URL.Path == "/user/login" {
		return
	}

	if len(tokenS) == 0 {
		ctx.Abort(http.StatusUnauthorized, "没认证")
		return
	}

	mobile, isPass := token.CheckToken(tokenS[0])
	if isPass {
		ctx.Abort(http.StatusUnauthorized, "认证过期")
		return
	}
	rsp, _ := service.Svc.UserService.CheckUserIsExisted(context.Background(), &proto.User{
		Mobile: mobile,
	})
	if rsp.Status == "500" {
		ctx.Abort(http.StatusBadRequest, "无该用户")
		return
	}
}
