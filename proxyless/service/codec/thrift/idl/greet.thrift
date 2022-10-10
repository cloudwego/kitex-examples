namespace go proxyless

struct HelloRequest {
    1: required string Message,
}

struct HelloResponse {
    1: required string Message,
}

service GreetService {
    HelloResponse SayHello1(1: HelloRequest request);
	HelloResponse SayHello2(1: HelloRequest request);
}
