package rpc

import (
	"context"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"douyin/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var userClient userservice.Client

func InitUserRpc() {
	c, err := userservice.NewClient("user",
		client.WithSuite(tracing.NewClientSuite()),
		client.WithHostPorts("0.0.0.0:8889"))
	if err != nil {
		panic(err)
	}

	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) (int64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}

	if resp.Resp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.Resp.StatusCode), *resp.Resp.StatusMsg)
	}
	return *resp.UserId, nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.BaseResp.StatusCode), *resp.BaseResp.StatusMsg)
	}
	return resp.UserId, nil
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.GetBaseResp().StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), *resp.BaseResp.StatusMsg)
	}
	return resp.Users, nil
}
