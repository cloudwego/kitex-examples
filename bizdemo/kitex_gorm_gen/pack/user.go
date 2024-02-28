/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pack

import (
	"strconv"

	"github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/kitex_gen/user"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/model/model"
)

// Users Convert orm_gen.User list to user_gorm.User list
func Users(models []*model.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to user_gorm.User
func User(model *model.User) *user.User {
	if model == nil {
		return nil
	}
	gender, err := strconv.ParseInt(model.Gender, 10, 64)
	if err != nil {
		return nil
	}
	return &user.User{
		UserId:    int64(model.ID),
		Name:      model.Name,
		Gender:    user.Gender(gender),
		Age:       int64(model.Age),
		Introduce: model.Introduce,
	}
}
