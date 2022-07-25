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