// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: admin/service/v1/i_organization.proto

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

const OperationOrganizationServiceCreateOrganization = "/admin.service.v1.OrganizationService/CreateOrganization"
const OperationOrganizationServiceDeleteOrganization = "/admin.service.v1.OrganizationService/DeleteOrganization"
const OperationOrganizationServiceGetOrganization = "/admin.service.v1.OrganizationService/GetOrganization"
const OperationOrganizationServiceListOrganization = "/admin.service.v1.OrganizationService/ListOrganization"
const OperationOrganizationServiceUpdateOrganization = "/admin.service.v1.OrganizationService/UpdateOrganization"

type OrganizationServiceHTTPServer interface {
	// CreateOrganization 创建部门
	CreateOrganization(context.Context, *v11.CreateOrganizationRequest) (*emptypb.Empty, error)
	// DeleteOrganization 删除部门
	DeleteOrganization(context.Context, *v11.DeleteOrganizationRequest) (*emptypb.Empty, error)
	// GetOrganization 查询部门详情
	GetOrganization(context.Context, *v11.GetOrganizationRequest) (*v11.Organization, error)
	// ListOrganization 查询部门列表
	ListOrganization(context.Context, *v1.PagingRequest) (*v11.ListOrganizationResponse, error)
	// UpdateOrganization 更新部门
	UpdateOrganization(context.Context, *v11.UpdateOrganizationRequest) (*emptypb.Empty, error)
}

func RegisterOrganizationServiceHTTPServer(s *http.Server, srv OrganizationServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/orgs", _OrganizationService_ListOrganization0_HTTP_Handler(srv))
	r.GET("/admin/v1/orgs/{id}", _OrganizationService_GetOrganization0_HTTP_Handler(srv))
	r.POST("/admin/v1/orgs", _OrganizationService_CreateOrganization0_HTTP_Handler(srv))
	r.PUT("/admin/v1/orgs/{org.id}", _OrganizationService_UpdateOrganization0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/orgs/{id}", _OrganizationService_DeleteOrganization0_HTTP_Handler(srv))
}

func _OrganizationService_ListOrganization0_HTTP_Handler(srv OrganizationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrganizationServiceListOrganization)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrganization(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListOrganizationResponse)
		return ctx.Result(200, reply)
	}
}

func _OrganizationService_GetOrganization0_HTTP_Handler(srv OrganizationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetOrganizationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrganizationServiceGetOrganization)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOrganization(ctx, req.(*v11.GetOrganizationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.Organization)
		return ctx.Result(200, reply)
	}
}

func _OrganizationService_CreateOrganization0_HTTP_Handler(srv OrganizationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateOrganizationRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrganizationServiceCreateOrganization)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateOrganization(ctx, req.(*v11.CreateOrganizationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _OrganizationService_UpdateOrganization0_HTTP_Handler(srv OrganizationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateOrganizationRequest
		if err := ctx.Bind(&in.Org); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrganizationServiceUpdateOrganization)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateOrganization(ctx, req.(*v11.UpdateOrganizationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _OrganizationService_DeleteOrganization0_HTTP_Handler(srv OrganizationServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteOrganizationRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationOrganizationServiceDeleteOrganization)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteOrganization(ctx, req.(*v11.DeleteOrganizationRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type OrganizationServiceHTTPClient interface {
	CreateOrganization(ctx context.Context, req *v11.CreateOrganizationRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteOrganization(ctx context.Context, req *v11.DeleteOrganizationRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetOrganization(ctx context.Context, req *v11.GetOrganizationRequest, opts ...http.CallOption) (rsp *v11.Organization, err error)
	ListOrganization(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListOrganizationResponse, err error)
	UpdateOrganization(ctx context.Context, req *v11.UpdateOrganizationRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type OrganizationServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewOrganizationServiceHTTPClient(client *http.Client) OrganizationServiceHTTPClient {
	return &OrganizationServiceHTTPClientImpl{client}
}

func (c *OrganizationServiceHTTPClientImpl) CreateOrganization(ctx context.Context, in *v11.CreateOrganizationRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/orgs"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrganizationServiceCreateOrganization))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrganizationServiceHTTPClientImpl) DeleteOrganization(ctx context.Context, in *v11.DeleteOrganizationRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/orgs/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrganizationServiceDeleteOrganization))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrganizationServiceHTTPClientImpl) GetOrganization(ctx context.Context, in *v11.GetOrganizationRequest, opts ...http.CallOption) (*v11.Organization, error) {
	var out v11.Organization
	pattern := "/admin/v1/orgs/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrganizationServiceGetOrganization))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrganizationServiceHTTPClientImpl) ListOrganization(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListOrganizationResponse, error) {
	var out v11.ListOrganizationResponse
	pattern := "/admin/v1/orgs"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationOrganizationServiceListOrganization))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *OrganizationServiceHTTPClientImpl) UpdateOrganization(ctx context.Context, in *v11.UpdateOrganizationRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/orgs/{org.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationOrganizationServiceUpdateOrganization))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Org, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
