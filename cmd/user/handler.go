package main

import (
	"context"
	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
)

type UserServiceImpl struct{}

func (u UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (r *user.CheckUserResponse, err error) {
	resp := user.NewCheckUserResponse()
	if len(req.GetUsername()) == 0 || len(req.GetPassword()) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	userID, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, err
	}

	resp.SetUserId(userID)
	resp.SetBaseResp(errno.BuildBaseResp(errno.Success))
	return resp, nil
}

func (u UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (r *user.MGetUserResponse, err error) {
	resp := user.NewMGetUserResponse()

	if len(req.UserIds) == 0 {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, err
	}

	resp.SetUsers(users)
	resp.SetBaseResp(errno.BuildBaseResp(errno.Success))
	return resp, nil
}

func (u UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (r *user.CreateUserResponse, err error) {
	resp := user.NewCreateUserResponse()

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp.Resp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userID, err := service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.Resp = errno.BuildBaseResp(err)
		return resp, err
	}

	resp.SetUserId(&userID)
	resp.SetResp(errno.BuildBaseResp(errno.Success))
	return resp, nil
}

var _ user.UserService = UserServiceImpl{}
