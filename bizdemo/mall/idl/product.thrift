include "base.thrift"
namespace go cmp.ecom.product

struct Brand {
    1: i64 BrandId,
    2: i64 ShopId,
    3: string BrandName,
    4: string Logo,
    5: string BrandStory
}

struct AddBrandReq {
    1: i64 ShopId,
    2: string BrandName,
    3: string Logo,
    4: string BrandStroy
}

struct AddBrandResp {
    1: i64 BrandId,
    255: base.BaseResp BaseResp
}

struct UpdateBrandReq {
    1: i64 BrandId,
    2: i64 ShopId,
    3: optional string BrandName,
    4: optional string Logo,
    5: optional string BrandStory
}

struct UpdateBrandResp {
    255: base.BaseResp BaseResp
}

struct DeleteBrandReq {
    1: i64 BrandId,
    2: i64 ShopId,
}

struct DeleteBrandResp {
    255: base.BaseResp BaseResp
}

struct GetBrandsByShopIdReq {
    1: i64 ShopId,
}

struct GetBrandsByShopIdResp {
    1: list<Brand> Brands,

    255: base.BaseResp BaseResp
}

service ProductService {
    AddBrandResp AddBrand(1: AddBrandReq req)
    UpdateBrandResp UpdateBrand(1: UpdateBrandReq req)
    DeleteBrandResp DeleteBrand(1: DeleteBrandReq req)
    GetBrandsByShopIdResp GetBrandsByShopId(1: GetBrandsByShopIdReq req)
}



