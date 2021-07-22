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

namespace go http

struct BizRequest {
    1: i64 vint64(api.query = 'vint64', api.vd = "$>0&&$<200")
    2: string text(api.body = 'text')
    3: i32 token(api.header = 'token')
    6: list<string> items(api.query = 'items')
    7: i32 version(api.path = 'version')
}

struct BizResponse {
    1: i32 token(api.header = 'token')
    2: string text(api.body = 'text')
    5: i32 http_code(api.http_code = '') 
}

service BizService {
    BizResponse BizMethod1(1: BizRequest req)(api.get = '/life/client/:version', api.baseurl = 'example.com', api.param = 'true')
    BizResponse BizMethod2(1: BizRequest req)(api.post = '/life/client/:version', api.baseurl = 'example.com', api.param = 'true', api.serializer = 'form')
    BizResponse BizMethod3(1: BizRequest req)(api.post = '/life/client/:version/other', api.baseurl = 'example.com', api.param = 'true', api.serializer = 'json')
}