package service

import (
	"context"
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/auth/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	service "github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/usecases/interfaces"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/utils"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/verification"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/anypb"
)

type AuthService struct {
	employeeUseCase service.EmployeeUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(userUseCase service.EmployeeUseCase) AuthService {
	return AuthService{employeeUseCase: userUseCase}
}

func (e *AuthService) PostSignup(ctx context.Context, req *pb.SignUpRequest) (*pb.Response, error) {

	var signup domain.Employee

	copier.Copy(&signup, &req)

	fmt.Println(signup)

	if err := e.employeeUseCase.SignUpOtp(ctx, signup); err == nil {
		return &pb.Response{
			Statuscode: 400,
			Message:    "User already exist",
		}, err
	}

	resp, err := verification.SendOtp(signup.Phone)

	if err != nil {
		return &pb.Response{
			Statuscode: 500,
			Message:    resp,
		}, err
	}

	token, err := utils.GenerateTokenForOtp(signup)

	if err != nil {
		return &pb.Response{
			Statuscode: 500,
			Message:    "Unable to signup",
		}, err
	}

	fmt.Println(token)

	return &pb.Response{
		Statuscode: 200,
		Message:    "Otp send succesfully",
		Data: &anypb.Any{
			Value: []byte(token),
		},
	}, nil
}

func (e *AuthService) VerifyOtp(ctx context.Context, req *pb.VerifyOtpRequest) (*pb.Response, error) {

	details, err := utils.ValidateOtpTokens(req.Cookie)

	if err != nil {
		return &pb.Response{
			Statuscode: 500,
			Message:    "Unable to find details",
		}, err
	}

	err = verification.ValidateOtp(details.Phone, req.Code)

	if err != nil {
		return &pb.Response{
			Statuscode: 400,
			Message:    "Invalid Otp",
		}, err
	}

	var signup domain.Employee
	copier.Copy(&signup, &details)

	err = e.employeeUseCase.SignUp(ctx, signup)

	if err != nil {
		return &pb.Response{
			Statuscode: 400,
			Message:    "Invalid Server Error",
		}, err
	}

	return &pb.Response{
		Statuscode: 200,
		Message:    "Successfully Account Created",
	}, nil

}

func (e *AuthService) PostLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	var login domain.Employee

	copier.Copy(&login, &req)

	emp, err := e.employeeUseCase.Login(ctx, login)

	if err != nil {
		return &pb.Response{
			Statuscode: 400,
			Message:    "Login failed",
		}, err
	}

	token, err := utils.GenerateTokens(emp.ID)

	fmt.Println(token)

	if err != nil {
		return &pb.Response{
			Statuscode: 500,
			Message:    "Unable to login",
		}, err
	}

	return &pb.Response{
		Statuscode: 200,
		Message:    "Signin succesfull",
		Data: &anypb.Any{
			Value: []byte(token),
		},
	}, nil
}
