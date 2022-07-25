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

enum Role {
    Admin,                  // 超级管理员
    CategoryOperator,       // 类目运营
    QualificationOperator,  // 资质运营
    ProductAuditor,         // 商品审核员
}

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

struct AddUserRoleReq {
    1: required string UserName,
    2: required Role Role,
}

struct AddUserRoleResp {
    255: base.BaseResp BaseResp
}

struct DelUserRoleReq {
    1: required string UserName,
    2: required Role Role,
}

struct DelUserRoleResp {
    255: base.BaseResp BaseResp
}

struct ValidateUserRolesReq {
    1: required string UserName,
    2: required list<Role> Roles,
}

struct ValidateUserRoleResp {
    1: bool IsPass,

    255: base.BaseResp BaseResp
}

service UserService {
    // 账户服务
    CreateUserResp CreateUser(1: CreateUserReq req)
    MGetUserResp MGetUser(1: MGetUserReq req)
    CheckUserResp CheckUser(1: CheckUserReq req)

    // 权限服务
    AddUserRoleResp AddUserRole(1: AddUserRoleReq req)
    DelUserRoleResp DelUserRole(1: DelUserRoleReq req)
    ValidateUserRoleResp ValidateUserRole(1: ValidateUserRolesReq req)
}
