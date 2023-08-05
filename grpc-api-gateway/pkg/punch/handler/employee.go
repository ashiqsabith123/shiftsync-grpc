package handler

import (
	"context"
	"net/http"
	"strconv"

	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/pb"
	"github.com/gin-gonic/gin"
)

func GetDuty(ctx *gin.Context, c pb.PunchServiceClient) {
	tempid, ok := ctx.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	res, _ := c.GetDuty(context.Background(), &pb.GetDutyRequest{
		Id: int32(id),
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(200, "duty schedules", res.Data)
	ctx.JSON(200, resp)

}

func PunchIn(ctx *gin.Context, c pb.PunchServiceClient) {
	tempid, ok := ctx.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	res, _ := c.PunchIn(context.Background(), &pb.PunchInRequest{
		Id: int32(id),
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(200, "Otp send to your verified phone number", nil)
	ctx.JSON(200, resp)
}

func VerifyOtpPunchin(ctx *gin.Context, c pb.PunchServiceClient) {
	var otp models.OtpStruct
	if err := ctx.ShouldBindJSON(&otp); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), otp)
		ctx.JSON(400, resp)
		return
	}

	tempid, ok := ctx.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	res, _ := c.VerifyOtpPunchin(context.Background(), &pb.PunchInRequestOtp{
		Code: otp.Code,
		Id:   int32(id),
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(200, "Punched succesfully", nil)
	ctx.JSON(200, resp)
}

func PunchOut(ctx *gin.Context, c pb.PunchServiceClient) {

	tempid, ok := ctx.Get("userId")
	id, _ := strconv.Atoi(tempid.(string))

	if !ok {
		resp := response.ErrorResponse(http.StatusInternalServerError, "employee id not found", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)
		return
	}

	res, _ := c.PunchOut(context.Background(), &pb.PunchOutRequest{
		Id: int32(id),
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)
		return
	}

	resp := response.SuccessResponse(200, "Punchout succesfully", nil)
	ctx.JSON(200, resp)
}
