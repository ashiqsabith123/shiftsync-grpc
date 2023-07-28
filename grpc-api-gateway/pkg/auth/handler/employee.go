package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth/models"
	"github.com/ashiqsabith123/shiftsync-grps-api-gateway/pkg/auth/pb"
	"github.com/gin-gonic/gin"
)

func PostLogin(ctx *gin.Context, c pb.AuthServiceClient) {
	var body models.SignUp

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(body)

	res, err := c.PostSignup(context.Background(), &pb.SignUpRequest{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Phone:     body.Phone,
		Email:     body.Email,
		Username:  body.UserName,
		Password:  body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Statuscode), res.Errors)
}
