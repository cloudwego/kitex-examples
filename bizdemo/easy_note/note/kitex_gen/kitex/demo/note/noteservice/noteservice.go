// Code generated by Kitex v0.1.4. DO NOT EDIT.

package noteservice

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/kitex_gen/kitex/demo/note"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return noteServiceServiceInfo
}

var noteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NoteService"
	handlerType := (*note.NoteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateNote": kitex.NewMethodInfo(createNoteHandler, newNoteServiceCreateNoteArgs, newNoteServiceCreateNoteResult, false),
		"MGetNote":   kitex.NewMethodInfo(mGetNoteHandler, newNoteServiceMGetNoteArgs, newNoteServiceMGetNoteResult, false),
		"DelNote":    kitex.NewMethodInfo(delNoteHandler, newNoteServiceDelNoteArgs, newNoteServiceDelNoteResult, false),
		"QueryNote":  kitex.NewMethodInfo(queryNoteHandler, newNoteServiceQueryNoteArgs, newNoteServiceQueryNoteResult, false),
		"UpdateNote": kitex.NewMethodInfo(updateNoteHandler, newNoteServiceUpdateNoteArgs, newNoteServiceUpdateNoteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "note",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.1.4",
		Extra:           extra,
	}
	return svcInfo
}

func createNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceCreateNoteArgs)
	realResult := result.(*note.NoteServiceCreateNoteResult)
	success, err := handler.(note.NoteService).CreateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceCreateNoteArgs() interface{} {
	return note.NewNoteServiceCreateNoteArgs()
}

func newNoteServiceCreateNoteResult() interface{} {
	return note.NewNoteServiceCreateNoteResult()
}

func mGetNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceMGetNoteArgs)
	realResult := result.(*note.NoteServiceMGetNoteResult)
	success, err := handler.(note.NoteService).MGetNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceMGetNoteArgs() interface{} {
	return note.NewNoteServiceMGetNoteArgs()
}

func newNoteServiceMGetNoteResult() interface{} {
	return note.NewNoteServiceMGetNoteResult()
}

func delNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceDelNoteArgs)
	realResult := result.(*note.NoteServiceDelNoteResult)
	success, err := handler.(note.NoteService).DelNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceDelNoteArgs() interface{} {
	return note.NewNoteServiceDelNoteArgs()
}

func newNoteServiceDelNoteResult() interface{} {
	return note.NewNoteServiceDelNoteResult()
}

func queryNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceQueryNoteArgs)
	realResult := result.(*note.NoteServiceQueryNoteResult)
	success, err := handler.(note.NoteService).QueryNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceQueryNoteArgs() interface{} {
	return note.NewNoteServiceQueryNoteArgs()
}

func newNoteServiceQueryNoteResult() interface{} {
	return note.NewNoteServiceQueryNoteResult()
}

func updateNoteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*note.NoteServiceUpdateNoteArgs)
	realResult := result.(*note.NoteServiceUpdateNoteResult)
	success, err := handler.(note.NoteService).UpdateNote(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newNoteServiceUpdateNoteArgs() interface{} {
	return note.NewNoteServiceUpdateNoteArgs()
}

func newNoteServiceUpdateNoteResult() interface{} {
	return note.NewNoteServiceUpdateNoteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateNote(ctx context.Context, req *note.CreateNoteRequest) (r *note.CreateNoteResponse, err error) {
	var _args note.NoteServiceCreateNoteArgs
	_args.Req = req
	var _result note.NoteServiceCreateNoteResult
	if err = p.c.Call(ctx, "CreateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetNote(ctx context.Context, req *note.MGetNoteRequest) (r *note.MGetNoteResponse, err error) {
	var _args note.NoteServiceMGetNoteArgs
	_args.Req = req
	var _result note.NoteServiceMGetNoteResult
	if err = p.c.Call(ctx, "MGetNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DelNote(ctx context.Context, req *note.DelNoteRequest) (r *note.DelNoteResponse, err error) {
	var _args note.NoteServiceDelNoteArgs
	_args.Req = req
	var _result note.NoteServiceDelNoteResult
	if err = p.c.Call(ctx, "DelNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryNote(ctx context.Context, req *note.QueryNoteRequest) (r *note.QueryNoteResponse, err error) {
	var _args note.NoteServiceQueryNoteArgs
	_args.Req = req
	var _result note.NoteServiceQueryNoteResult
	if err = p.c.Call(ctx, "QueryNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateNote(ctx context.Context, req *note.UpdateNoteRequest) (r *note.UpdateNoteResponse, err error) {
	var _args note.NoteServiceUpdateNoteArgs
	_args.Req = req
	var _result note.NoteServiceUpdateNoteResult
	if err = p.c.Call(ctx, "UpdateNote", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
