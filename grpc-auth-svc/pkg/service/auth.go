package service

import (
	"context"
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/pb"
	service "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases/interfaces"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/utils"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/verification"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/anypb"
)

type EmployeeHandler struct {
	employeeUseCase service.EmployeeUseCase
	pb.UnimplementedAuthServiceServer
}

func NewEmployeeHandler(userUseCase service.EmployeeUseCase) EmployeeHandler {
	return EmployeeHandler{employeeUseCase: userUseCase}
}

func (e *EmployeeHandler) PostSignup(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {

	var signup domain.Employee

	copier.Copy(&signup, &req)

	fmt.Println(signup)

	if err := e.employeeUseCase.SignUpOtp(ctx, signup); err == nil {
		// resp := response.ErrorResponse(400, "User already exist", "", nil)
		// ctxt.JSON(http.StatusBadRequest, resp)
		return &pb.SignUpResponse{
			Statuscode: 400,
			Message:    "User already exist",
		}, err
	}

	resp, err := verification.SendOtp(signup.Phone)

	if err != nil {
		return &pb.SignUpResponse{
			Statuscode: 500,
			Message:    resp,
		}, err
	}

	token, err := utils.GenerateTokenForOtp(signup)

	if err != nil {
		return &pb.SignUpResponse{
			Statuscode: 500,
			Message:    "Unable to signup",
		}, err
	}

	fmt.Println(token)

	return &pb.SignUpResponse{
		Statuscode: 200,
		Message:    "Otp send succesfully",
		Data: &anypb.Any{
			Value: []byte(token),
		},
	}, nil
}
