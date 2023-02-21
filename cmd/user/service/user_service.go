package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/infra/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/snowid"
	"fmt"
	"io"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) CreateUser(req *user.CreateUserRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	userID, err := snowid.GenID()
	if err != nil {
		return 0, err
	}
	u := &db.User{
		UserID:   userID,
		Username: req.GetUsername(),
		Password: MD5(req.GetPassword()),
	}
	if err = db.CreateUser(s.ctx, u); err != nil {
		return 0, err
	}
	return u.UserID, nil
}

func (s *UserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.GetUsername())
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}

	u := users[0]
	if u.Password != MD5(req.Password) {
		return 0, errno.LoginErr
	}
	return u.UserID, nil
}

func (s *UserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	users, err := db.MGetUser(s.ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}

	ret := make([]*user.User, 0)
	for _, DBModel := range users {
		userTar := &user.User{
			Id:            DBModel.UserID,
			Name:          DBModel.Username,
			FollowCount:   &DBModel.FollowCount,
			FollowerCount: &DBModel.FollowerCount,
			// TODO: call relation rpc
			IsFollow: false,
		}
		ret = append(ret, userTar)
	}
	return ret, nil
}

func MD5(v string) string {
	h := md5.New()
	io.WriteString(h, v)
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return passWord
}
