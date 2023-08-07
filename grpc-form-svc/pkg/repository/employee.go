package repository

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/encrypt"
	repo "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/repository/interface"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type employeeDatabase struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) repo.EmployeeRepository {
	return &employeeDatabase{DB: DB}
}

func (e *employeeDatabase) CheckFormDetails(cntxt context.Context, form domain.Form) (domain.Form, bool) {

	fmt.Println("b", form)
	var details domain.Form
	if err := e.DB.Where("form_id = ? OR account_no = ? OR pan_number = ? OR adhaar_no = ? ", form.FormID, base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Account_no))), base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number))), base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Adhaar_no)))).First(&details).Error; err != nil {

		return details, false
	}

	return details, true
}

func (e *employeeDatabase) AddForm(cntxt context.Context, form domain.Form) error {

	if err := e.DB.Create(&form).Error; err != nil {
		return err
	}

	return nil
}

func (e *employeeDatabase) FormCorrection(ctx context.Context, form domain.Form) error {
	var formDetails domain.Form
	if err := e.DB.Where("form_id = ?", form.FormID).First(&formDetails).Error; err != nil {
		return err
	}

	copier.Copy(&formDetails, &form)

	e.DB.Save(&formDetails)
	return nil
}
