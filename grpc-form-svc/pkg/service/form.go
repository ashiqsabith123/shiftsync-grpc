package service

import (
	"context"

	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/pb"
	service "github.com/ashiqsabith123/shiftsync-grpc-form-svc/pkg/usecases/interface"
	"github.com/jinzhu/copier"
)

type FormService struct {
	employeeUseCase service.EmployeeUseCase
	pb.UnimplementedFormServiceServer
}

func NewFormService(employeeusecase service.EmployeeUseCase) FormService {
	return FormService{
		employeeUseCase: employeeusecase,
	}
}

func (e *FormService) PostForm(ctx context.Context, req *pb.FormRequest) (*pb.Response, error) {

	var form domain.Form

	form.FormID = int(req.Id)
	copier.Copy(&form, &req)

	if err := e.employeeUseCase.AddForm(ctx, form); err != nil {
		if err != nil {
			return &pb.Response{
				Statuscode: 400,
				Message:    "Eroor when adding form",
			}, err
		}

	}

	return &pb.Response{
		Statuscode: 200,
		Message:    "Form submitted succesfully pending for verification",
	}, nil

}
