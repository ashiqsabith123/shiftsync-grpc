package repository

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/domain"
)

type EmployeeRepository interface {
	CheckFormDetails(cntxt context.Context, form domain.Form) (domain.Form, bool)
	AddForm(cntxt context.Context, form domain.Form) error
	FormCorrection(ctx context.Context, form domain.Form) error
}
