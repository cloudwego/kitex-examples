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

package main

import (
	"log"
	pb "pb_generic_streaming_demo/kitex_gen/pb/streamingservice"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := pb.NewServer(new(StreamingServiceImpl),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
