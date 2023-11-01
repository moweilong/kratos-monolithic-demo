// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: admin/service/v1/i_menu.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pagination "kratos-monolithic-demo/gen/api/go/common/pagination"
	v1 "kratos-monolithic-demo/gen/api/go/system/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenuServiceCreateMenu = "/admin.service.v1.MenuService/CreateMenu"
const OperationMenuServiceDeleteMenu = "/admin.service.v1.MenuService/DeleteMenu"
const OperationMenuServiceGetMenu = "/admin.service.v1.MenuService/GetMenu"
const OperationMenuServiceListMenu = "/admin.service.v1.MenuService/ListMenu"
const OperationMenuServiceUpdateMenu = "/admin.service.v1.MenuService/UpdateMenu"

type MenuServiceHTTPServer interface {
	// CreateMenu 创建菜单
	CreateMenu(context.Context, *v1.CreateMenuRequest) (*v1.Menu, error)
	// DeleteMenu 删除菜单
	DeleteMenu(context.Context, *v1.DeleteMenuRequest) (*emptypb.Empty, error)
	// GetMenu 查询菜单详情
	GetMenu(context.Context, *v1.GetMenuRequest) (*v1.Menu, error)
	// ListMenu 查询菜单列表
	ListMenu(context.Context, *pagination.PagingRequest) (*v1.ListMenuResponse, error)
	// UpdateMenu 更新菜单
	UpdateMenu(context.Context, *v1.UpdateMenuRequest) (*v1.Menu, error)
}

func RegisterMenuServiceHTTPServer(s *http.Server, srv MenuServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/menus", _MenuService_ListMenu0_HTTP_Handler(srv))
	r.GET("/admin/v1/menus/{id}", _MenuService_GetMenu0_HTTP_Handler(srv))
	r.POST("/admin/v1/menus", _MenuService_CreateMenu0_HTTP_Handler(srv))
	r.PUT("/admin/v1/menus/{id}", _MenuService_UpdateMenu0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/menus/{id}", _MenuService_DeleteMenu0_HTTP_Handler(srv))
}

func _MenuService_ListMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in pagination.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceListMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*pagination.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.ListMenuResponse)
		return ctx.Result(200, reply)
	}
}

func _MenuService_GetMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceGetMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenu(ctx, req.(*v1.GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_CreateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.CreateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*v1.CreateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_UpdateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.UpdateMenuRequest
		if err := ctx.Bind(&in.Menu); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceUpdateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenu(ctx, req.(*v1.UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_DeleteMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.DeleteMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*v1.DeleteMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type MenuServiceHTTPClient interface {
	CreateMenu(ctx context.Context, req *v1.CreateMenuRequest, opts ...http.CallOption) (rsp *v1.Menu, err error)
	DeleteMenu(ctx context.Context, req *v1.DeleteMenuRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetMenu(ctx context.Context, req *v1.GetMenuRequest, opts ...http.CallOption) (rsp *v1.Menu, err error)
	ListMenu(ctx context.Context, req *pagination.PagingRequest, opts ...http.CallOption) (rsp *v1.ListMenuResponse, err error)
	UpdateMenu(ctx context.Context, req *v1.UpdateMenuRequest, opts ...http.CallOption) (rsp *v1.Menu, err error)
}

type MenuServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuServiceHTTPClient(client *http.Client) MenuServiceHTTPClient {
	return &MenuServiceHTTPClientImpl{client}
}

func (c *MenuServiceHTTPClientImpl) CreateMenu(ctx context.Context, in *v1.CreateMenuRequest, opts ...http.CallOption) (*v1.Menu, error) {
	var out v1.Menu
	pattern := "/admin/v1/menus"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuServiceCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) DeleteMenu(ctx context.Context, in *v1.DeleteMenuRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/menus/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) GetMenu(ctx context.Context, in *v1.GetMenuRequest, opts ...http.CallOption) (*v1.Menu, error) {
	var out v1.Menu
	pattern := "/admin/v1/menus/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceGetMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) ListMenu(ctx context.Context, in *pagination.PagingRequest, opts ...http.CallOption) (*v1.ListMenuResponse, error) {
	var out v1.ListMenuResponse
	pattern := "/admin/v1/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceListMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) UpdateMenu(ctx context.Context, in *v1.UpdateMenuRequest, opts ...http.CallOption) (*v1.Menu, error) {
	var out v1.Menu
	pattern := "/admin/v1/menus/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuServiceUpdateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Menu, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
