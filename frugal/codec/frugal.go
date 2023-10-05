// Copyright 2023 CloudWeGo Authors
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
	"fmt"
	"reflect"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/frugal"
	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
)

type Meta struct {
	MethodName  string
	MessageType thrift.TMessageType
	SeqID       int32
}

func HasFrugalTag(data interface{}) bool {
	dt := reflect.TypeOf(data).Elem()
	if dt.NumField() > 0 && dt.Field(0).Tag.Get("frugal") == "" {
		return false
	}
	return true
}

func FrugalEncode(data interface{}, meta Meta) ([]byte, error) {
	if !HasFrugalTag(data) {
		return nil, fmt.Errorf("frugal tag missing")
	}
	// calculate and malloc message buffer
	msgBeginLen := bthrift.Binary.MessageBeginLength(meta.MethodName, meta.MessageType, meta.SeqID)
	msgEndLen := bthrift.Binary.MessageEndLength()
	objectLen := frugal.EncodedSize(data)
	buf := make([]byte, msgBeginLen+objectLen+msgEndLen)

	// encode message
	offset := bthrift.Binary.WriteMessageBegin(buf, meta.MethodName, meta.MessageType, meta.SeqID)
	writeLen, err := frugal.EncodeObject(buf[offset:], nil, data)
	if err != nil {
		return buf, fmt.Errorf("thrift marshal, frugal.EncodeObject failed: %s", err.Error())
	}
	offset += writeLen
	bthrift.Binary.WriteMessageEnd(buf[offset:])
	return buf, nil
}

func FrugalDecode(buf []byte, data interface{}) (Meta, error) {
	if !HasFrugalTag(data) {
		return Meta{}, fmt.Errorf("frugal tag missing")
	}
	methodName, messageType, seqID, length, err := bthrift.Binary.ReadMessageBegin(buf)
	if err != nil {
		return Meta{}, err
	}
	_, err = frugal.DecodeObject(buf[length:], data)
	return Meta{
		MethodName:  methodName,
		MessageType: messageType,
		SeqID:       seqID,
	}, err
}

func main() {
	meta := Meta{
		MethodName:  "echo",
		MessageType: thrift.CALL,
		SeqID:       0x01020304,
	}
	req := &api.EchoEchoArgs{
		Req: &api.Request{
			Message: "Hello",
		},
	}
	buf, err := FrugalEncode(req, meta)
	fmt.Printf("Encode: buf = %v, err = %v\n", buf, err)

	decodedReq := &api.EchoEchoArgs{}
	meta, err = FrugalDecode(buf, decodedReq)
	fmt.Printf("Decode: req = %v, meta = %v err = %v\n", decodedReq, meta, err)
}
