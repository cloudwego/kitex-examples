// Code generated by Kitex v0.0.6. DO NOT EDIT.
package echo

import (
	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler pbapi.Echo, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
