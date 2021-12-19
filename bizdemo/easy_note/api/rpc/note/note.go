package note

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/opentracing/opentracing-go"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/note/noteservice"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var (
	noteClient noteservice.Client
)

func initJaeger(service string) (client.Suite, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return trace.NewDefaultClientSuite(), closer
}

func Init() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	tracer, _ := initJaeger(constant.ServiceName)

	fp := retry.NewFailurePolicy()

	c, err := noteservice.NewClient(
		constant.NoteServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    //mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		client.WithFailureRetry(fp),                    // retry
		client.WithSuite(tracer),                       // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}

	noteClient = c

}

func CreateNote(ctx context.Context, req *note.CreateNoteRequest) error {
	resp, err := noteClient.CreateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func MGetNotes(ctx context.Context, req *note.MGetNoteRequest) ([]*note.Note, error) {
	resp, err := noteClient.MGetNote(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Notes, nil
}

func QueryNotes(ctx context.Context, req *note.QueryNoteRequest) ([]*note.Note, int64, error) {
	resp, err := noteClient.QueryNote(ctx, req)
	if err != nil {
		return nil, 0, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Notes, resp.Total, nil
}

func UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) error {
	resp, err := noteClient.UpdateNote(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func DelNote(ctx context.Context, req *note.DelNoteRequest) error {
	resp, err := noteClient.DelNote(ctx, req)
	if err != nil {
		return err
	}

	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}
