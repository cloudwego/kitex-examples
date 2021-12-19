package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
)

func MGetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	conn := GetDBReader(ctx)
	if err := conn.Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func CreateUser(ctx context.Context, users []*model.User) error {
	conn := GetDBWriter(ctx)
	return conn.Create(users).Error
}

func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	conn := GetDBReader(ctx)
	res := make([]*model.User, 0)
	if err := conn.Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
