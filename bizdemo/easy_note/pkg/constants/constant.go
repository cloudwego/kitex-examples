// Copyright 2021 CloudWeGo Authors
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

package constants

import (
	"os"
)

const (
	NoteTableName            = "note"
	UserTableName            = "user"
	IdentityKey              = "id"
	Total                    = "total"
	Notes                    = "notes"
	NoteID                   = "note_id"
	ApiServiceName           = "demoapi"
	NoteServiceName          = "demonote"
	UserServiceName          = "demouser"
	CPURateLimit     float64 = 80.0
	DefaultLimit             = 10
	defaultSecretKey         = "secret key"
)

func SecretKey() string {
	if key := os.Getenv("JWT_SECRET_KEY"); key != "" {
		return key
	}
	return defaultSecretKey
}

func MySQLDSN() string {
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		return dsn
	}
	return MySQLDefaultDSN
}

var (
	MySQLDefaultDSN = "gorm:gorm@tcp(" + GetIp("MysqlIp") + ":9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = GetIp("EtcdIp") + ":2379"
)
