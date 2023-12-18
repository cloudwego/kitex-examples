// idl/hello.proto
namespace go hello.example

struct HelloReq {
    1: string Name;
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request);
}
