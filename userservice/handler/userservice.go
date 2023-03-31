package handler

import (
	"context"
	"database/sql"
	"fmt"
	"go-micro.dev/v4/logger"
	"golang.org/x/crypto/bcrypt"
	"time"
	cache "userservice/db/redisInit"
	db "userservice/db/sqlc"
	pb "userservice/proto"
	"userservice/utils"
)

type UserService struct {
}

func (u UserService) SendCaptcha(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	_, err := db.DB.GetUserByPhone(ctx, sql.NullString{String: user.GetMobile(), Valid: true})
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，发送验证码失败,因为 %s", err),
			response)
		return nil
	}
	utils.SendMessage(user.Mobile)
	return nil
}

func (u UserService) LoginByMobileCaptcha(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	us, err := db.DB.GetUserByPhone(ctx, sql.NullString{String: user.GetMobile(), Valid: true})
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，查密码失败,因为 %s", err),
			response)
		return nil
	}

	//pwd := user.Password
	valid, err := cache.IsValidCaptcha(us.Mobile.String, user.Captcha)
	if !valid {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，验证码错误,因为 %s", err),
			response)
		return nil
	}
	db.DB.UpdateUserLoginDate(context.Background(), db.UpdateUserLoginDateParams{LoginDate: sql.NullTime{Time: time.Now(), Valid: true}, Mobile: sql.NullString{
		String: user.Mobile,
		Valid:  true,
	}})
	utils.PackInfo(err,
		"100",
		fmt.Sprintf("验证码正确"),
		response)
	return nil
}

func (u UserService) CheckUserIsExisted(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	_, err := db.DB.GetUserByPhone(ctx, sql.NullString{String: user.GetMobile(), Valid: true})
	if err != nil {
		response.Description = err.Error()
		response.Status = "500"
		return nil
	}
	response.Description = ""
	response.Status = "100"
	return nil
}

func (u UserService) LoginByMobile(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	us, err := db.DB.GetUserByPhone(ctx, sql.NullString{String: user.GetMobile(), Valid: true})
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，查密码失败,因为 %s", err),
			response)
		return nil
	}

	//pwd := user.Password
	err = bcrypt.CompareHashAndPassword([]byte(us.Password), []byte(user.Password))
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("密码错误"),
			response)
		return nil
	}
	db.DB.UpdateUserLoginDate(context.Background(), db.UpdateUserLoginDateParams{LoginDate: sql.NullTime{Time: time.Now(), Valid: true}, Mobile: sql.NullString{
		String: user.Mobile,
		Valid:  true,
	}})
	utils.PackInfo(err,
		"100",
		fmt.Sprintf("密码正确"),
		response)
	return nil
}

func (u UserService) GetUser(ctx context.Context, request *pb.User, user *pb.User) error {
	us, err := db.DB.GetUserByID(ctx, request.GetId())
	if err != nil {
		logger.Errorf("GetUser is error of %s", err)
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，获取用户失败,因为 %s", err),
			user.Response)
		return nil
	}
	fmt.Println(*user)
	//fmt.Println(us)
	user.Address = us.Address.String
	user.Id = us.ID
	user.Mobile = us.Mobile.String
	user.Avatar = us.Avatar.String
	user.Email = us.Email.String
	user.Name = us.Name.String
	user.Username = us.Username
	user.Birth = &pb.Date{Time: us.Birth.Time.String()}
	return nil
}

func (u UserService) GetUserByPhone(ctx context.Context, request *pb.User, user *pb.User) error {
	//fmt.Println(request)
	us, err := db.DB.GetUserByPhone(ctx, sql.NullString{String: request.Mobile, Valid: true})
	if err != nil {
		logger.Errorf("GetUserPhone is error of %s", err)

		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，使用电话获取用户失败,因为 %s", err),
			user.Response)

		return nil
	}
	fmt.Println(us.Email.String)
	user.Address = us.Address.String
	user.Id = us.ID
	user.Mobile = us.Mobile.String
	user.Avatar = us.Avatar.String
	user.Email = us.Email.String
	user.Name = us.Name.String
	user.Username = us.Username
	user.Birth = &pb.Date{Time: us.Birth.Time.String()}
	user.Response = new(pb.UserResponse)
	fmt.Println(user)
	utils.PackInfo(nil,
		"100",
		fmt.Sprintf("操作成功"),
		user.Response)
	return nil
}

func (u UserService) RegisterUser(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	registerParams := db.RegisterUserParams{
		Username: user.Username,
		Password: utils.HashAndSalt([]byte(user.Password)),
		Name:     sql.NullString{String: user.Name, Valid: true},
		Email:    sql.NullString{String: user.Email, Valid: true},
		Mobile:   sql.NullString{String: user.Mobile, Valid: true},
		Gender:   sql.NullInt32{Int32: user.Gender, Valid: true},
	}
	res, err := db.DB.RegisterUser(ctx, registerParams)
	if err != nil {
		logger.Errorf("RegisterUser is error of %s", err)
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，注册用户失败,因为 %s", err),
			response)
		return nil
	}
	num, err := res.RowsAffected()
	if num == 1 {
		utils.PackInfo(err,
			"100",
			fmt.Sprintf("操作成功，创建用户成功"),
			response)
	} else {

		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，注册用户失败,因为 %s", err),
			response)
	}
	return nil
}

func (u UserService) ChangePasswordToUser(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	updateUserPasswordParams := db.UpdateUserPasswordParams{
		Password: user.Password,
		Username: user.Username,
		ID:       user.Id,
	}
	err := db.DB.UpdateUserPassword(ctx, updateUserPasswordParams)
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，更新密码失败,因为 %s", err),
			response)
	} else {

		utils.PackInfo(err,
			"100",
			fmt.Sprintf("操作成功"),
			response)
	}
	return nil
}

func (u UserService) CheckPasswordToUser(ctx context.Context, request *pb.User, response *pb.UserResponse) error {
	user, err := db.DB.GetUserByName(ctx, request.GetUsername())
	dBUserPassword := user.Password
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，查密码失败,因为 %s", err),
			response)
		return nil
	}
	pwd := utils.HashAndSalt([]byte(request.Password))
	if dBUserPassword != pwd {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("密码不正确"),
			response)
		return nil
	}
	utils.PackInfo(err,
		"100",
		fmt.Sprintf("密码正确"),
		response)
	return nil
}

func (u UserService) ChangeInfoToUser(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	updateUserInfoParams := db.UpdateUserInfoParams{
		Avatar:   sql.NullString{String: user.Avatar, Valid: true},
		Username: user.Username,
		Name:     sql.NullString{String: user.Name, Valid: true},
		Gender:   sql.NullInt32{Int32: user.Gender, Valid: true},
		ID:       user.Id,
	}
	err := db.DB.UpdateUserInfo(ctx, updateUserInfoParams)
	if err != nil {
		utils.PackInfo(err,
			"500",
			fmt.Sprintf("操作失败，更新信息失败,因为 %s", "密码不对"),
			response)

		return nil
	}
	utils.PackInfo(err,
		"100",
		fmt.Sprintf("操作正确"),
		response)

	return nil
}

func (u UserService) UnregisterUser(ctx context.Context, user *pb.User, response *pb.UserResponse) error {
	//TODO implement me
	panic("implement me")
}
