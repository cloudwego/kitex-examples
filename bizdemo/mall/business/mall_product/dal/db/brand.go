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

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"gorm.io/gorm"
)

type BrandDO struct {
	gorm.Model
	ShopId     int64  `json:"shop_id"`
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	BrandStory string `json:"brand_story"`
}

func (d *BrandDO) TableName() string {
	return conf.BrandTableName
}

func GetBrandInfoByShopId(ctx context.Context, shopId int64) ([]*BrandDO, error) {
	brandList := make([]*BrandDO, 0)
	if err := DB.WithContext(ctx).Where("shop_id = ?", shopId).Find(&brandList).Error; err != nil {
		return nil, err
	}
	return brandList, nil
}

func CreateBrand(ctx context.Context, brand *BrandDO) (int64, error) {
	if err := DB.WithContext(ctx).Create(brand).Error; err != nil {
		return 0, err
	}
	return int64(brand.ID), nil
}

func DeleteBrand(ctx context.Context, brandId, shopId int64) error {
	return DB.WithContext(ctx).Where("id = ? and shop_id = ?", brandId, shopId).Delete(&BrandDO{}).Error
}

func UpdateBrand(ctx context.Context, brandId, shopId int64, name, logo, brandStory *string) error {
	updateMap := map[string]interface{}{}
	if name != nil {
		updateMap["name"] = *name
	}
	if logo != nil {
		updateMap["logo"] = *logo
	}
	if brandStory != nil {
		updateMap["brand_story"] = *brandStory
	}

	return DB.WithContext(ctx).Model(&BrandDO{}).Where("id = ? and shop_id = ?", brandId, shopId).
		Updates(updateMap).Error
}
