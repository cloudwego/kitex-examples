package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
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

func (s *UserService) CreateUser(req *user.CreateUserReq) error {
	users, err := db.QueryUser(s.ctx, req.GetUserName())
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}

func (s *UserService) MGetUser(req *user.MGetUserReq) ([]*user.User, error) {
	users, err := db.MGetUsers(s.ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	ret := make([]*user.User, 0)
	for _, userModel := range users {
		ret = append(ret, &user.User{
			UserId:   int64(userModel.ID),
			UserName: userModel.UserName,
		})
	}
	return ret, nil
}

func (s *UserService) CheckUser(req *user.CheckUserReq) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserName
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}
