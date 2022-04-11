// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
include "base.thrift"
namespace go cmp.ecom.user

struct User {
    1: i64 UserId,
    2: string UserName
}

struct CreateUserReq {
    1: string UserName,
    2: string Password
}

struct CreateUserResp {
    255: base.BaseResp BaseResp
}

struct MGetUserReq {
    1: list<i64> Ids,
}

struct MGetUserResp {
    1: list<User> Users,

    255: base.BaseResp BaseResp
}

struct CheckUserReq {
    1: string UserName,
    2: string Password
}

struct CheckUserResp {
    1: i64 UserId,

    255: base.BaseResp BaseResp
}

service UserService {
    CreateUserResp CreateUser(1: CreateUserReq req)
    MGetUserResp MGetUser(1: MGetUserReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)
}
