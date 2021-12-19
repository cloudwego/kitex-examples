package pack

import (
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
)

func User(u *model.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{UserId: int64(u.ID), UserName: u.UserName, Avatar: "test"}
}

func Users(us []*model.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
