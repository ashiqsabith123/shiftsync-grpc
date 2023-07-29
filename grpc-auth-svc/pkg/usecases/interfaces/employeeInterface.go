package interfaces

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
)

type EmployeeUseCase interface {
	SignUpOtp(r context.Context, find domain.Employee) error
}
