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
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/ent"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/kitex_gen/user"
)

// Users Convert model.User list to user_gorm.User list
func Users(models []*ent.User) []*user.User {
	users := make([]*user.User, 0, len(models))
	for _, m := range models {
		if u := User(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// User Convert model.User to user_gorm.User
func User(model *ent.User) *user.User {
	if model == nil {
		return nil
	}
	return &user.User{
		UserId:    int64(model.ID),
		Name:      model.Name,
		Gender:    user.Gender(model.Gender),
		Age:       int64(model.Age),
		Introduce: model.Introduce,
	}
}
