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

package conf

const (
	UserTableName              = "t_user"
	UserRoleTableName          = "t_user_role"
	ShopTableName              = "t_shop"
	BrandTableName             = "t_brand"
	CategoryTableName          = "t_category"
	AttributeTableName         = "t_attribute"
	ShopCategoryTableName      = "t_shop_category"
	CategoryAttributeTableName = "t_category_attribute"

	SecretKey   = "secret key"
	IdentityKey = "id"

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"

	UserRpcServiceName    = "cmp.ecom.user"
	ShopRpcServiceName    = "cmp.ecom.shop"
	ProductRpcServiceName = "cmp.ecom.product"
)
