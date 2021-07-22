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

package main

import (
	"context"
	"log"

	"github.com/bytedance/gopkg/cloud/circuitbreaker"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
)

func changeHandler(key string, oldState, newState circuitbreaker.State, m circuitbreaker.Metricer) {
	log.Printf("circuitbreaker status change, old: %v, new: %v\n", oldState, newState)
}

func getKey(ctx context.Context, request interface{}) (key string, enabled bool) {
	return "1234", true
}

func getErrorType(ctx context.Context, request, response interface{}, err error) circuitbreak.ErrorType {
	if err != nil {
		return circuitbreak.TypeFailure
	}
	return circuitbreak.TypeSuccess
}

func decorateError(ctx context.Context, request interface{}, err error) error {
	return err
}
