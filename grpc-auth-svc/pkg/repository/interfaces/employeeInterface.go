package interfaces

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
)

type EmployeeRepository interface {
	FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error)
}
