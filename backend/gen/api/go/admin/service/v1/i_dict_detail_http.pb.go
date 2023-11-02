// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/service/v1/i_dict_detail.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDictDetailServiceCreateDictDetail = "/admin.service.v1.DictDetailService/CreateDictDetail"
const OperationDictDetailServiceDeleteDictDetail = "/admin.service.v1.DictDetailService/DeleteDictDetail"
const OperationDictDetailServiceGetDictDetail = "/admin.service.v1.DictDetailService/GetDictDetail"
const OperationDictDetailServiceListDictDetail = "/admin.service.v1.DictDetailService/ListDictDetail"
const OperationDictDetailServiceUpdateDictDetail = "/admin.service.v1.DictDetailService/UpdateDictDetail"

type DictDetailServiceHTTPServer interface {
	// CreateDictDetail 创建字典详情
	CreateDictDetail(context.Context, *v11.CreateDictDetailRequest) (*v11.DictDetail, error)
	// DeleteDictDetail 删除字典详情
	DeleteDictDetail(context.Context, *v11.DeleteDictDetailRequest) (*emptypb.Empty, error)
	// GetDictDetail 查询字典详情
	GetDictDetail(context.Context, *v11.GetDictDetailRequest) (*v11.DictDetail, error)
	// ListDictDetail 查询字典详情列表
	ListDictDetail(context.Context, *v1.PagingRequest) (*v11.ListDictDetailResponse, error)
	// UpdateDictDetail 更新字典详情
	UpdateDictDetail(context.Context, *v11.UpdateDictDetailRequest) (*v11.DictDetail, error)
}

func RegisterDictDetailServiceHTTPServer(s *http.Server, srv DictDetailServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/dict:details", _DictDetailService_ListDictDetail0_HTTP_Handler(srv))
	r.GET("/admin/v1/dict:details/{id}", _DictDetailService_GetDictDetail0_HTTP_Handler(srv))
	r.POST("/admin/v1/dict:details", _DictDetailService_CreateDictDetail0_HTTP_Handler(srv))
	r.PUT("/admin/v1/dict:details/{id}", _DictDetailService_UpdateDictDetail0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/dict:details/{id}", _DictDetailService_DeleteDictDetail0_HTTP_Handler(srv))
}

func _DictDetailService_ListDictDetail0_HTTP_Handler(srv DictDetailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDetailServiceListDictDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDictDetail(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListDictDetailResponse)
		return ctx.Result(200, reply)
	}
}

func _DictDetailService_GetDictDetail0_HTTP_Handler(srv DictDetailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetDictDetailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDetailServiceGetDictDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDictDetail(ctx, req.(*v11.GetDictDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.DictDetail)
		return ctx.Result(200, reply)
	}
}

func _DictDetailService_CreateDictDetail0_HTTP_Handler(srv DictDetailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateDictDetailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDetailServiceCreateDictDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDictDetail(ctx, req.(*v11.CreateDictDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.DictDetail)
		return ctx.Result(200, reply)
	}
}

func _DictDetailService_UpdateDictDetail0_HTTP_Handler(srv DictDetailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateDictDetailRequest
		if err := ctx.Bind(&in.Detail); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDetailServiceUpdateDictDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDictDetail(ctx, req.(*v11.UpdateDictDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.DictDetail)
		return ctx.Result(200, reply)
	}
}

func _DictDetailService_DeleteDictDetail0_HTTP_Handler(srv DictDetailServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteDictDetailRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDetailServiceDeleteDictDetail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDictDetail(ctx, req.(*v11.DeleteDictDetailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type DictDetailServiceHTTPClient interface {
	CreateDictDetail(ctx context.Context, req *v11.CreateDictDetailRequest, opts ...http.CallOption) (rsp *v11.DictDetail, err error)
	DeleteDictDetail(ctx context.Context, req *v11.DeleteDictDetailRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetDictDetail(ctx context.Context, req *v11.GetDictDetailRequest, opts ...http.CallOption) (rsp *v11.DictDetail, err error)
	ListDictDetail(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListDictDetailResponse, err error)
	UpdateDictDetail(ctx context.Context, req *v11.UpdateDictDetailRequest, opts ...http.CallOption) (rsp *v11.DictDetail, err error)
}

type DictDetailServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDictDetailServiceHTTPClient(client *http.Client) DictDetailServiceHTTPClient {
	return &DictDetailServiceHTTPClientImpl{client}
}

func (c *DictDetailServiceHTTPClientImpl) CreateDictDetail(ctx context.Context, in *v11.CreateDictDetailRequest, opts ...http.CallOption) (*v11.DictDetail, error) {
	var out v11.DictDetail
	pattern := "/admin/v1/dict:details"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictDetailServiceCreateDictDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDetailServiceHTTPClientImpl) DeleteDictDetail(ctx context.Context, in *v11.DeleteDictDetailRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/dict:details/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDetailServiceDeleteDictDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDetailServiceHTTPClientImpl) GetDictDetail(ctx context.Context, in *v11.GetDictDetailRequest, opts ...http.CallOption) (*v11.DictDetail, error) {
	var out v11.DictDetail
	pattern := "/admin/v1/dict:details/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDetailServiceGetDictDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDetailServiceHTTPClientImpl) ListDictDetail(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListDictDetailResponse, error) {
	var out v11.ListDictDetailResponse
	pattern := "/admin/v1/dict:details"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDetailServiceListDictDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDetailServiceHTTPClientImpl) UpdateDictDetail(ctx context.Context, in *v11.UpdateDictDetailRequest, opts ...http.CallOption) (*v11.DictDetail, error) {
	var out v11.DictDetail
	pattern := "/admin/v1/dict:details/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictDetailServiceUpdateDictDetail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Detail, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
