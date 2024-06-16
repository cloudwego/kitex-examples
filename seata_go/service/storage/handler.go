// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"

	"github.com/cloudwego/kitex-examples/seata_go/service/storage/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/storage/dal/mysql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// StorageServiceImpl implements the last service interface defined in the IDL.
type StorageServiceImpl struct{}

// Deduct implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) Deduct(ctx context.Context, commodityCode string, count int32) (err error) {
	err = mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		account, err := model.Get(ctx, tx, commodityCode)
		if err != nil {
			klog.Errorf("get storage failed: %v", err)
			return err
		}
		klog.Infof("%#v", account)
		err = model.Update(ctx, tx, commodityCode, account.Count-count)
		if err != nil {
			klog.Errorf("update storage failed: %v", err)
			return err
		}
		return nil
	})
	return
}

// Calculate implements the StorageServiceImpl interface.
func (s *StorageServiceImpl) Calculate(ctx context.Context, commodityCode string, count int32) (resp int32, err error) {
	commodity, err := model.Get(ctx, mysql.DB, commodityCode)
	if err != nil {
		klog.Errorf("get storage failed: %v", err)
		return 0, err
	}
	return commodity.Price * count, nil
}
