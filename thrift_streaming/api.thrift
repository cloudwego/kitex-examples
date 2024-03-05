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

namespace go echo

struct Request {
    1: required string message,
}

struct Response {
    1: required string message,
}

service TestService {
    Response Echo (1: Request req) (streaming.mode="bidirectional"),
    Response EchoClient (1: Request req) (streaming.mode="client"),
    Response EchoServer (1: Request req) (streaming.mode="server"),
    Response EchoUnary (1: Request req) (streaming.mode="unary"), // not recommended

    Response EchoPingPong (1: Request req), // KitexThrift, non-streaming
}
