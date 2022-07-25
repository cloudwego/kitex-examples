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
namespace go cmp.ecom.governance

struct AsyncAuditProductReq {
    1: i64 ProductId,
    2: string Name, // 商品名
    3: string Description, // 详情
    4: string img, // 主图
    5: string RecommendRemark, // 推荐语
    6: list<string> SpecPics, // 规格图
}

struct AsyncAuditProductResp {
    255: base.BaseResp BaseResp
}

struct GetProductAuditRecordsReq {
    1: i64 ProductId,
}

enum AuditStatus {
    Pending, // 审核中
    Pass, // 审核通过
    Reject, // 审核驳回
}

struct AuditRecord {
    1: string AuditTime,
    2: AuditStatus AuditStatus,
    3: string RejectReason,
}

struct GetProductAuditRecordsResp {
    1: list<AuditRecord> Records,

    255: base.BaseResp BaseResp
}

service GovernanceService {
    AsyncAuditProductResp AsyncAuditProduct(1: AsyncAuditProductReq req) // 异步提审
    GetProductAuditRecordsResp GetProductAuditRecords(1: GetProductAuditRecordsReq req) // 查询审核记录
}