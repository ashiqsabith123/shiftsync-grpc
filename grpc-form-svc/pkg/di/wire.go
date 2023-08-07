//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/db"
	repo "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/repository"
	service "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/service"
	usecase "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/usecases"
	"github.com/google/wire"
)

func InitializeApi(config config.Config) service.FormService {
	wire.Build(
		db.ConnectDatabase,
		repo.NewEmployeeRepository,
		usecase.NewEmployeeUseCase,
		service.NewFormService,
	)

	return service.FormService{}
}
