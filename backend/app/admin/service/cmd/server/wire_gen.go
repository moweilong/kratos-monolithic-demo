// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-bootstrap/gen/api/go/conf/v1"
	"kratos-monolithic-demo/app/admin/service/internal/data"
	"kratos-monolithic-demo/app/admin/service/internal/server"
	"kratos-monolithic-demo/app/admin/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	authenticator := data.NewAuthenticator(bootstrap)
	engine := data.NewAuthorizer()
	entClient := data.NewEntClient(bootstrap, logger)
	client := data.NewRedisClient(bootstrap, logger)
	dataData, cleanup, err := data.NewData(entClient, client, authenticator, engine, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userTokenRepo := data.NewUserTokenRepo(dataData, authenticator, logger)
	authenticationService := service.NewAuthenticationService(logger, userRepo, userTokenRepo)
	userService := service.NewUserService(logger, userRepo)
	dictRepo := data.NewDictRepo(dataData, logger)
	dictService := service.NewDictService(logger, dictRepo)
	dictDetailRepo := data.NewDictDetailRepo(dataData, logger)
	dictDetailService := service.NewDictDetailService(logger, dictDetailRepo)
	menuRepo := data.NewMenuRepo(dataData, logger)
	menuService := service.NewMenuService(menuRepo, logger)
	routerService := service.NewRouterService(logger, menuRepo)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	organizationService := service.NewOrganizationService(organizationRepo, logger)
	roleRepo := data.NewRoleRepo(dataData, logger)
	roleService := service.NewRoleService(roleRepo, logger)
	positionRepo := data.NewPositionRepo(dataData, logger)
	positionService := service.NewPositionService(positionRepo, logger)
	httpServer := server.NewRESTServer(bootstrap, logger, authenticator, engine, authenticationService, userService, dictService, dictDetailService, menuService, routerService, organizationService, roleService, positionService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
