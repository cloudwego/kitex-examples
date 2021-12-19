package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {

	h := md5.New()
	io.WriteString(h, req.Password)
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
