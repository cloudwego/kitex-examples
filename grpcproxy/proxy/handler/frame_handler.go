// Copyright 2022 CloudWeGo Authors
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

package handler

import (
	"context"
	"io"
	"log"

	"github.com/cloudwego/kitex-examples/grpcproxy/proxy"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/netpoll"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

// grpc proxy connection pool
var connPool = nphttp2.NewConnPool("", 2, grpc.ConnectOptions{
	KeepaliveParams:       grpc.ClientKeepalive{},
	InitialWindowSize:     0,
	InitialConnWindowSize: 0,
	WriteBufferSize:       0,
	ReadBufferSize:        0,
	MaxHeaderListSize:     nil,
	ShortConn:             false,
})

// GRPCFrameProxyHandler redirects grpc frame from client to target server.
func GRPCFrameProxyHandler(ctx context.Context, methodName string, stream streaming.Stream) error {
	log.Println("Proxy Handler is working....")

	// find target address by methodName
	network, address := proxy.Resolve(methodName)

	// create a new RPC Info and modify some infos if you want.
	sri := rpcinfo.GetRPCInfo(ctx)
	ri := rpcinfo.NewRPCInfo(sri.From(), sri.To(), sri.Invocation(), sri.Config(), sri.Stats())
	clientCtx := rpcinfo.NewCtxWithRPCInfo(context.Background(), ri)

	conn, err := connPool.Get(clientCtx, network, address, remote.ConnOption{
		Dialer:         netpoll.NewDialer(),
		ConnectTimeout: 0,
	})
	if err != nil {
		return err
	}
	clientConn := conn.(nphttp2.GRPCConn)
	defer func() {
		clientConn.Close()
		connPool.Put(clientConn)
	}()

	serverConn, err := nphttp2.GetServerConn(stream)
	if err != nil {
		return err
	}

	s2c := redirectFrame(serverConn, clientConn)
	c2s := redirectFrame(clientConn, serverConn)

	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2c:
			if s2cErr != io.EOF {
				return s2cErr
			}
		case c2sErr := <-c2s:
			if c2sErr != io.EOF {
				return c2sErr
			}
		}
	}
	return nil
}

func redirectFrame(from, to nphttp2.GRPCConn) chan error {
	ret := make(chan error)

	go func() {
		for {
			hdr, data, err := from.ReadFrame()
			if err != nil {
				// write last empty data frame with END_STREAM flag
				to.WriteFrame(hdr, data)
				ret <- err
				break
			}
			_, err = to.WriteFrame(hdr, data)
			if err != nil {
				ret <- err
				break
			}
		}
	}()

	return ret
}
