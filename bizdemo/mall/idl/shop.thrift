include "base.thrift"
namespace go cmp.ecom.shop

struct SettleShopReq {
    1: string UserName,
}

struct SettleShopResp {
    1: i64 ShopId,

    255: base.BaseResp BaseResp
}

struct GetShopIdByNameReq {
    1: string UserName,
}

struct GetShopIdByNameResp {
    1: i64 ShopId,

    255: base.BaseResp BaseResp
}

service ShopService {
    SettleShopResp SettleShop(1: SettleShopReq req)
    GetShopIdByNameResp GetShopIdByName(1: GetShopIdByNameReq req)
}