//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/db"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/service"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases"
	"github.com/google/wire"
)

func InitializeApi(config config.Config) service.EmployeeHandler {
	wire.Build(
		db.ConnectDatabase,
		repository.NewEmployeeRepository,
		usecases.NewEmployeeUseCase,
		service.NewEmployeeHandler,
	)
	return service.EmployeeHandler{}

}
