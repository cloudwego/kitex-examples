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
