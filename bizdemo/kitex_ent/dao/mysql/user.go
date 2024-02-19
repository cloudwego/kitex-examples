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

package mysql

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/ent"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/ent/user"
)

func CreateUser(client *ent.Client, name string, age int) (*ent.User, error) {
	return client.User.
		Create().
		SetName(name).
		SetAge(age).
		Save(context.Background())
}

func GetUser(client *ent.Client, id int) (*ent.User, error) {
	return client.User.Get(context.Background(), id)
}

func UpdateUser(client *ent.Client, id int, name string, age int) error {
	_, err := client.User.
		Update().
		Where(user.ID(id)).
		SetName(name).
		SetAge(age).
		Save(context.Background())
	return err
}

func SearchUsers(client *ent.Client, keyword string, page, pageSize int) ([]*ent.User, error) {
	offset := (page - 1) * pageSize
	return client.User.
		Query().
		Where(user.NameContains(keyword)).
		Offset(offset).
		Limit(pageSize).All(context.Background())
}

func DeleteUser(client *ent.Client, id int) error {
	return client.User.DeleteOneID(id).Exec(context.Background())
}
