// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/db"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/service"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases"
)

// Injectors from wire.go:

func InitializeApi(config2 config.Config) service.EmployeeHandler {
	gormDB := db.ConnectDatabase(config2)
	employeeRepository := repository.NewEmployeeRepository(gormDB)
	employeeUseCase := usecases.NewEmployeeUseCase(employeeRepository)
	employeeHandler := service.NewEmployeeHandler(employeeUseCase)
	return employeeHandler
}