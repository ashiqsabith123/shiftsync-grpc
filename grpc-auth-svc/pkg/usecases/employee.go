package usecases

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository/interfaces"
	service "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases/interfaces"
)

type employeeUseCase struct {
	employeeRepo repo.EmployeeRepository
}

func NewEmployeeUseCase(rep repo.EmployeeRepository) service.EmployeeUseCase {
	return &employeeUseCase{employeeRepo: rep}
}

func (u *employeeUseCase) SignUpOtp(r context.Context, find domain.Employee) error {
	_, err := u.employeeRepo.FindEmployee(r, find)

	return err
}
