include "base.thrift"
namespace go cmp.ecom.shop

struct SettleShopReq {
    1: i64 UserId,
    2: string ShopName,
}

struct SettleShopResp {
    1: i64 ShopId,

    255: base.BaseResp BaseResp
}

struct GetShopIdByUserIdReq {
    1: i64 UserId,
}

struct GetShopIdByUserIdResp {
    1: i64 ShopId,

    255: base.BaseResp BaseResp
}

service ShopService {
    SettleShopResp SettleShop(1: SettleShopReq req)
    GetShopIdByUserIdResp GetShopIdByUserId(1: GetShopIdByUserIdReq req)
}