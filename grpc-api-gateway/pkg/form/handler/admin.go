package handler

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
	"github.com/gin-gonic/gin"
)

func ViewApplications(ctx *gin.Context, c pb.FormServiceClient) {
	res, err := c.ViewApplications(ctx, &pb.ViewApplicationRequest{})

	if err != nil || len(res.Responce) == 0 {
		ctx.JSON(404, gin.H{
			"status":  404,
			"message": "no new forms found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"forms":  res.Responce,
	})

}

func ApproveApplication(ctx *gin.Context, c pb.FormServiceClient) {

	var id models.FormApproveID
	if err := ctx.ShouldBindJSON(&id); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), id)
		ctx.JSON(400, resp)
		return
	}
	res, _ := c.ApproveApplication(ctx, &pb.ApproveApplicationRequest{
		Id: int32(id.FormID),
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)

		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Form Approved Succesfully", nil)
	ctx.JSON(resp.StatusCode, resp)

}

func ApplicationCorrection(ctx *gin.Context, c pb.FormServiceClient) {

	var corDetails models.FormCorrection
	if err := ctx.ShouldBindJSON(&corDetails); err != nil {
		resp := response.ErrorResponse(400, "invalid input", err.Error(), corDetails)
		ctx.JSON(400, resp)
		return
	}

	res, _ := c.ApplicationCorrection(ctx, &pb.ApplicationCorrectionRequest{
		Formid:     int32(corDetails.FormID),
		Correction: corDetails.Correction,
	})

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)

		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Correction Submited Succesfully", nil)
	ctx.JSON(resp.StatusCode, resp)

}
