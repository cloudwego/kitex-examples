package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}

	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	io.WriteString(h, req.Password)
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*model.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}
