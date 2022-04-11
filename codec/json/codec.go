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

package json

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	"github.com/cloudwego/kitex/pkg/remote/codec/perrors"
)

type JsonCodec struct {
	printDebugInfo bool
}

func NewJsonCodec(debug bool) remote.Codec {
	return &JsonCodec{printDebugInfo: debug}
}

func (jc *JsonCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	var validData interface{}
	switch message.MessageType() {
	case remote.Exception:
		switch e := message.Data().(type) {
		case *remote.TransError:
			validData = &Exception{e.TypeID(), e.Error()}
		case error:
			validData = &Exception{remote.InternalError, e.Error()}
		default:
			return errors.New("exception relay must implement error type")
		}
	default:
		validData = message.Data()
	}
	payload, err := json.Marshal(validData)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, marshal payload failed: %w", err))
	}
	if jc.printDebugInfo {
		klog.Infof("encoded payload: %s\n", payload)
	}
	data := &Meta{
		ServiceName: message.RPCInfo().Invocation().ServiceName(),
		MethodName:  message.RPCInfo().Invocation().MethodName(),
		SeqID:       message.RPCInfo().Invocation().SeqID(),
		MsgType:     uint32(message.MessageType()),
		Payload:     payload,
	}
	buf, err := json.Marshal(data)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, marshal data failed: %w", err))
	}
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(len(buf)))
	_, err = out.WriteBinary(length)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, write length failed: %w", err))
	}
	_, err = out.WriteBinary(buf)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json encode, write data failed: %w", err))
	}
	return nil
}

func (jc *JsonCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	length, err := in.ReadBinary(8)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, read length failed: %w", err))
	}
	buf, err := in.ReadBinary(int(binary.BigEndian.Uint64(length)))
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, read data failed: %w", err))
	}
	data := &Meta{}
	err = json.Unmarshal(buf, data)
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal data failed: %w", err))
	}
	if err = codec.SetOrCheckSeqID(data.SeqID, message); err != nil {
		return err
	}
	if err = codec.SetOrCheckMethodName(data.MethodName, message); err != nil {
		return err
	}
	if err = codec.NewDataIfNeeded(data.MethodName, message); err != nil {
		return err
	}
	if jc.printDebugInfo {
		klog.Infof("encoded payload: %s\n", data.Payload)
	}
	if remote.MessageType(data.MsgType) == remote.Exception {
		var exception Exception
		err = json.Unmarshal(data.Payload, &exception)
		if err != nil {
			return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal payload failed: %w", err))
		}
		return exception
	}
	err = json.Unmarshal(data.Payload, message.Data())
	if err != nil {
		return perrors.NewProtocolError(fmt.Errorf("json decode, unmarshal payload failed: %w", err))
	}
	return nil
}

func (jc *JsonCodec) Name() string {
	return "JSON"
}
