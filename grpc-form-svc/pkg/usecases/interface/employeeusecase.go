package usecases

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/domain"
)

type EmployeeUseCase interface {
	AddForm(ctx context.Context, form domain.Form) error
}
