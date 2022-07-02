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

package db

import "gorm.io/gorm"

type (
	Product struct {
		gorm.Model
		ProductId        int64  `json:"product_id"`        // 商品ID
		ShopId           int64  `json:"shop_id"`           // 店铺ID
		Name             string `json:"name"`              // 商品名
		Description      string `json:"description"`       // 详情
		Img              string `json:"img"`               // 主图
		Status           int64  `json:"status"`            // 商品状态
		AuditStatus      int64  `json:"audit_status"`      // 审核状态
		CategoryId       int64  `json:"category_id"`       // 类目id
		CategoryProperty string `json:"category_property"` // 类目属性 json键值对
		BrandId          int64  `json:"brand_id"`          // 品牌id
		RecommendRemark  string `json:"recommend_remark"`  // 推荐语
		CommitTime       string `json:"commit_time"`       // 提交时间
		RejectReason     string `json:"reject_reason"`     // 驳回原因
		Extra            string `json:"extra"`
	}
)
