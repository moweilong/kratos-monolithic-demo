// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: admin/service/v1/i_position.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-monolithic-demo/api/gen/go/user/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPositionServiceCreatePosition = "/admin.service.v1.PositionService/CreatePosition"
const OperationPositionServiceDeletePosition = "/admin.service.v1.PositionService/DeletePosition"
const OperationPositionServiceGetPosition = "/admin.service.v1.PositionService/GetPosition"
const OperationPositionServiceListPosition = "/admin.service.v1.PositionService/ListPosition"
const OperationPositionServiceUpdatePosition = "/admin.service.v1.PositionService/UpdatePosition"

type PositionServiceHTTPServer interface {
	// CreatePosition 创建职位
	CreatePosition(context.Context, *v11.CreatePositionRequest) (*emptypb.Empty, error)
	// DeletePosition 删除职位
	DeletePosition(context.Context, *v11.DeletePositionRequest) (*emptypb.Empty, error)
	// GetPosition 查询职位详情
	GetPosition(context.Context, *v11.GetPositionRequest) (*v11.Position, error)
	// ListPosition 查询职位列表
	ListPosition(context.Context, *v1.PagingRequest) (*v11.ListPositionResponse, error)
	// UpdatePosition 更新职位
	UpdatePosition(context.Context, *v11.UpdatePositionRequest) (*emptypb.Empty, error)
}

func RegisterPositionServiceHTTPServer(s *http.Server, srv PositionServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/positions", _PositionService_ListPosition0_HTTP_Handler(srv))
	r.GET("/admin/v1/positions/{id}", _PositionService_GetPosition0_HTTP_Handler(srv))
	r.POST("/admin/v1/positions", _PositionService_CreatePosition0_HTTP_Handler(srv))
	r.PUT("/admin/v1/positions/{position.id}", _PositionService_UpdatePosition0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/positions/{id}", _PositionService_DeletePosition0_HTTP_Handler(srv))
}

func _PositionService_ListPosition0_HTTP_Handler(srv PositionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPositionServiceListPosition)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPosition(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListPositionResponse)
		return ctx.Result(200, reply)
	}
}

func _PositionService_GetPosition0_HTTP_Handler(srv PositionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetPositionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPositionServiceGetPosition)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPosition(ctx, req.(*v11.GetPositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Position)
		return ctx.Result(200, reply)
	}
}

func _PositionService_CreatePosition0_HTTP_Handler(srv PositionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreatePositionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPositionServiceCreatePosition)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePosition(ctx, req.(*v11.CreatePositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _PositionService_UpdatePosition0_HTTP_Handler(srv PositionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdatePositionRequest
		if err := ctx.Bind(&in.Position); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPositionServiceUpdatePosition)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePosition(ctx, req.(*v11.UpdatePositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _PositionService_DeletePosition0_HTTP_Handler(srv PositionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeletePositionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPositionServiceDeletePosition)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePosition(ctx, req.(*v11.DeletePositionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type PositionServiceHTTPClient interface {
	CreatePosition(ctx context.Context, req *v11.CreatePositionRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeletePosition(ctx context.Context, req *v11.DeletePositionRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetPosition(ctx context.Context, req *v11.GetPositionRequest, opts ...http.CallOption) (rsp *v11.Position, err error)
	ListPosition(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListPositionResponse, err error)
	UpdatePosition(ctx context.Context, req *v11.UpdatePositionRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type PositionServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPositionServiceHTTPClient(client *http.Client) PositionServiceHTTPClient {
	return &PositionServiceHTTPClientImpl{client}
}

func (c *PositionServiceHTTPClientImpl) CreatePosition(ctx context.Context, in *v11.CreatePositionRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/positions"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPositionServiceCreatePosition))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PositionServiceHTTPClientImpl) DeletePosition(ctx context.Context, in *v11.DeletePositionRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/positions/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPositionServiceDeletePosition))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PositionServiceHTTPClientImpl) GetPosition(ctx context.Context, in *v11.GetPositionRequest, opts ...http.CallOption) (*v11.Position, error) {
	var out v11.Position
	pattern := "/admin/v1/positions/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPositionServiceGetPosition))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PositionServiceHTTPClientImpl) ListPosition(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListPositionResponse, error) {
	var out v11.ListPositionResponse
	pattern := "/admin/v1/positions"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPositionServiceListPosition))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PositionServiceHTTPClientImpl) UpdatePosition(ctx context.Context, in *v11.UpdatePositionRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/positions/{position.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPositionServiceUpdatePosition))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Position, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
