package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/pack"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/service"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	resp = user.NewCreateUserResp()

	if len(req.GetUserName()) == 0 || len(req.GetPassword()) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserReq) (resp *user.MGetUserResp, err error) {
	resp = user.NewMGetUserResp()

	if len(req.Ids) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Users = users
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	resp = user.NewCheckUserResp()

	if len(req.Password) == 0 || len(req.UserName) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userId, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userId
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
