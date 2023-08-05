package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/pb"
	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
	"github.com/gin-gonic/gin"
)

func PostSignup(ctx *gin.Context, c pb.AuthServiceClient) {
	var body models.SignUp

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, _ := c.PostSignup(context.Background(), &pb.SignUpRequest{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Phone:     body.Phone,
		Email:     body.Email,
		Username:  body.UserName,
		Password:  body.Password,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	ctx.SetCookie("_signup-cookie", string(res.Data.Value), 20*60, "", "", false, true)

	resp := response.SuccessResponse(200, "Otp send succesfully", nil)
	ctx.JSON(200, resp)
}

func VerifyOtp(ctx *gin.Context, c pb.AuthServiceClient) {

	var body models.OtpStruct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, resp)
		return

	}

	token, err := ctx.Cookie("_signup-cookie")

	if err != nil {
		resp := response.ErrorResponse(500, "unable to find details", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	ctx.SetCookie("_signup-cookie", "", -1, "", "", false, true)

	res, _ := c.VerifyOtp(context.Background(), &pb.VerifyOtpRequest{
		Code:   body.Code,
		Cookie: token,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Succesfully account created", nil)
	ctx.JSON(resp.StatusCode, resp)
}

func PostLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	var values models.LoginStruct
	if err := ctx.ShouldBindJSON(&values); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), models.LoginStruct{})
		ctx.JSON(400, resp)
		return
	}

	if values.Username == "" || values.Password == "" {
		err := errors.New("missing credentials")
		resp := response.ErrorResponse(400, "Username and password is mandatory", err.Error(), models.LoginStruct{})
		ctx.JSON(400, resp)
		return
	}

	res, err := c.PostLogin(ctx, &pb.LoginRequest{
		Username: values.Username,
		Password: values.Password,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, err.Error(), nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	ctx.SetCookie("_employee-cookie", string(res.Data.Value), 20*60, "", "", false, true)

	resp := response.SuccessResponse(200, res.Message, nil)
	ctx.JSON(200, resp)
}
