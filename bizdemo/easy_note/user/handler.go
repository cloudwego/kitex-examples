package main

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CreateUserResponse)
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = errno.Success.ToBaseResp()
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.MGetUserResponse)
	if len(req.UserIds) == 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.Success.ToBaseResp()
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.CheckUserResponse)
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = errno.ParamErr.ToBaseResp()
		return resp, nil
	}
	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = errno.Success.ToBaseResp()
	return resp, nil
}
