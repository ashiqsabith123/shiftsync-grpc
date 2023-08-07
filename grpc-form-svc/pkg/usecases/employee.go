package usecases

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/encrypt"
	repo "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/repository/interface"
	service "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/usecases/interface"
)

type employeeUseCase struct {
	employeeRepo repo.EmployeeRepository
}

func NewEmployeeUseCase(rep repo.EmployeeRepository) service.EmployeeUseCase {
	return &employeeUseCase{employeeRepo: rep}
}

func (u *employeeUseCase) AddForm(ctx context.Context, form domain.Form) error {

	form.Account_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Account_no)))
	form.Pan_number = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Pan_number)))
	form.Adhaar_no = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Adhaar_no)))
	form.Ifsc_code = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Ifsc_code)))
	form.Name_as_per_passbokk = base64.StdEncoding.EncodeToString(encrypt.Encrypt([]byte(form.Name_as_per_passbokk)))
	form.Status = "P"

	details, ok := u.employeeRepo.CheckFormDetails(ctx, form)

	if ok && details.Status == "A" || details.Status == "P" {
		return errors.New("details already found")
	} else if ok && details.Status == "C" {
		if err := u.employeeRepo.FormCorrection(ctx, form); err != nil {
			return err
		}
	} else {
		if err := u.employeeRepo.AddForm(ctx, form); err != nil {
			return err
		}
	}

	return nil
}
