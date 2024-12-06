namespace cpp base
namespace py base
namespace go base
namespace java base

struct TrafficEnv {
    1: bool Open = false,
    2: string Env = "",
}

struct Base {
    1: string LogID = "",
    2: string Caller = "",
    3: string Addr = "",
    4: string client = "",
    5: optional TrafficEnv trafficEnv,
    6: optional map<string, string> extra,
}

struct BaseResp {
    1: string StatusMessage = "",
    2: i32 StatusCode = 0,
    3: optional map<string, string> Extra,
}
