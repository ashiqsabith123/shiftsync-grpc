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

	if err := e.DB.Where("id= ? OR email = ? OR phone = ? OR username = ?", find.ID, find.Email, find.Phone, find.Username).First(&emp).Error; err != nil {

		return find, errors.New("no user found")
	}

	cntxt.Done()
	return emp, nil
}

func (e *employeeDatabase) AddEmployee(cntxt context.Context, emp domain.Employee) error {

	if err := e.DB.Create(&emp).Error; err != nil {
		return errors.Join(errors.New("error from here"), err)
	}

	//err := e.DB.Raw("INSERT INTO employees (first_name, last_name, email, user_name, pass_word, phone) VALUES (?, ?, ?, ?, ?, ?) RETURNING id", emp.First_name, emp.Last_name, emp.Email, emp.User_name, emp.Pass_word, emp.Phone).Error
	return nil
}
