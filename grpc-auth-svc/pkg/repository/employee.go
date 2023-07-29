package repository

import (
	"context"
	"errors"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	repo "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) repo.EmployeeRepository {
	return &employeeDatabase{DB: DB}
}

func (e *employeeDatabase) FindEmployee(cntxt context.Context, find domain.Employee) (domain.Employee, error) {

	var emp domain.Employee

	if err := e.DB.Where("id= ? OR email = ? OR phone = ? OR user_name = ?", find.ID, find.Email, find.Phone, find.Username).First(&emp).Error; err != nil {

		return find, errors.New("no user found")
	}

	cntxt.Done()
	return emp, nil
}
