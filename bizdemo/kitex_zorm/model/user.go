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

package model

import "gitee.com/chunanyong/zorm"

type User struct {
	zorm.EntityStruct
	Id        int64  `json:"id" column:"id"`
	Name      string `json:"name" column:"name"`
	Gender    int64  `json:"gender" column:"gender"`
	Age       int64  `json:"age" column:"age"`
	Introduce string `json:"introduce" column:"introduce"`
}

func (entity *User) GetTableName() string {
	return "user"
}

func (entity *User) GetPKColumnName() string {
	return "id"
}
