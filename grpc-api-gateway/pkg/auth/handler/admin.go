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

func AdminPostLogin(ctx *gin.Context, c pb.AuthServiceClient) {
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

	res, _ := c.AdminPostLogin(ctx, &pb.LoginRequest{
		Username: values.Username,
		Password: values.Password,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, res.Errors, nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	//fmt.Println(res.Data.Value)

	ctx.SetCookie("_admin-cookie", string(res.Data.Value), 20*60, "", "", false, true)

	resp := response.SuccessResponse(200, res.Message, string(res.Data.Value))
	ctx.JSON(200, resp)
}

func AdminPostSignup(ctx *gin.Context, c pb.AuthServiceClient) {

	var rep models.AdminSignUp

	if err := ctx.ShouldBindJSON(&rep); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), rep)
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	res, _ := c.AdminPostSignup(context.Background(), &pb.SignUpRequest{
		Firstname: rep.FirstName,
		Lastname:  rep.LastName,
		Email:     rep.Email,
		Phone:     rep.Phone,
		Username:  rep.UserName,
		Password:  rep.Password,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Succesfully account created", nil)
	ctx.JSON(resp.StatusCode, resp)

}
