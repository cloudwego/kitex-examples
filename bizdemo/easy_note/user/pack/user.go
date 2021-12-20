// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
