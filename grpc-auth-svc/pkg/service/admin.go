package service

import (
	"context"
	"fmt"

	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/auth/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/domain"
	"github.com/ashiqsabith123/shiftsync-grpc-auth-svc/pkg/utils"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/anypb"
)

func (e *AuthService) AdminPostSignup(ctx context.Context, req *pb.SignUpRequest) (*pb.Response, error) {
	var adminDetails domain.Admin

	copier.Copy(&adminDetails, req)

	if err := e.adminUseCase.SignUp(ctx, adminDetails); err != nil {
		return &pb.Response{

			Statuscode: 400,
			Message:    err.Error(),
		}, nil
	}

	return &pb.Response{
		Statuscode: 201,
		Message:    "Suucesfully account created",
	}, nil
}

func (e *AuthService) AdminPostLogin(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {

	var login domain.Admin

	copier.Copy(&login, &req)

	emp, err := e.adminUseCase.SignIn(ctx, login)

	if err != nil {
		return &pb.Response{
			Statuscode: 400,
			Message:    "Login failed",
			Errors:     err.Error(),
		}, nil
	}

	token, err := utils.GenerateTokens(uint(emp.ID))

	fmt.Println(token)

	if err != nil {
		return &pb.Response{
			Statuscode: 500,
			Message:    "Unable to login",
		}, nil
	}

	return &pb.Response{
		Statuscode: 200,
		Message:    "Signin succesfull",
		Data: &anypb.Any{
			Value: []byte(token),
		},
	}, nil
}
