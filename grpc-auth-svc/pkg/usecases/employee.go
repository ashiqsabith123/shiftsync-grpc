package usecases

import (
	"context"
	"errors"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository/interfaces"
	service "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases/interfaces"
	"golang.org/x/crypto/bcrypt"
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

func (u *employeeUseCase) SignUp(cntxt context.Context, signup domain.Employee) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 14)
	if err != nil {
		return errors.New("bcrypt failed:" + err.Error())
	}

	signup.Password = string(hash)
	return u.employeeRepo.AddEmployee(cntxt, signup)
}

func (u *employeeUseCase) Login(r context.Context, login domain.Employee) (domain.Employee, error) {
	employee, err := u.employeeRepo.FindEmployee(r, login)

	if err != nil {
		//fmt.Println("Hello")
		return login, errors.New(err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(login.Password)); err != nil {

		return login, errors.New("incorrect password")
	}

	return employee, nil
}
