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

	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/mysql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// AccountServiceImpl implements the last service interface defined in the IDL.
type AccountServiceImpl struct{}

// Deduct implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) Deduct(ctx context.Context, userId string, money int32) (err error) {
	err = mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		account, err := model.Get(ctx, tx, userId)
		if err != nil {
			klog.Errorf("get account failed: %v", err)
			return err
		}
		err = model.Update(ctx, tx, userId, account.Money-money)
		if err != nil {
			klog.Errorf("update account failed: %v", err)
			return err
		}
		return nil
	})
	return
}
