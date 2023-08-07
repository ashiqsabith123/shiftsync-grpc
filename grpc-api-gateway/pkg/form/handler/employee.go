// package handler

// import (
// 	"context"
// 	"net/http"
// 	"strconv"

// 	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
// 	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
// 	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/copier"
// )

// func PostForm(ctx *gin.Context, c pb.FormServiceClient) {
// 	var tempForm models.Form

// 	empID, ok := ctx.Get("userId")

// 	if !ok || empID == "" {
// 		resp := response.ErrorResponse(500, "Employee id not found in cookie", "", nil)
// 		ctx.JSON(http.StatusInternalServerError, resp)
// 		return
// 	}

// 	if err := ctx.ShouldBindJSON(&tempForm); err != nil {
// 		resp := response.ErrorResponse(400, "Invalid input", err.Error(), tempForm)
// 		ctx.JSON(400, resp)
// 		return
// 	}

// 	var form pb.FormRequest

// 	tempid, _ := strconv.Atoi(empID.(string))

// 	form.Id = int32(tempid)
// 	copier.Copy(&form, &tempForm)

// 	res, _ := c.PostForm(context.Background(), &form)

// 	if res.Statuscode >= 400 {
// 		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
// 		ctx.JSON(resp.StatusCode, resp)
// 		return
// 	}

// 	resp := response.SuccessResponse(int(res.Statuscode), "Form submitted succesfully pending for verification", nil)
// 	ctx.JSON(resp.StatusCode, resp)

// }

package handler

import (
	"context"
	"net/http"
	"strconv"

	// Replace with the correct import path for your protobuf package

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func PostForm(ctx *gin.Context, c pb.FormServiceClient) {
	// Create a new logrus.Logger instance
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	var tempForm models.Form

	empID, ok := ctx.Get("userId")

	if !ok || empID == "" {
		resp := response.ErrorResponse(500, "Employee id not found in cookie", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)

		// Log the error
		log.WithFields(logrus.Fields{
			"microservice": "API Gateway",
			"endpoint":     ctx.FullPath(),
			"error":        "Employee id not found in cookie",
		}).Error("Error processing request")

		return
	}

	if err := ctx.ShouldBindJSON(&tempForm); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), tempForm)
		ctx.JSON(400, resp)

		// Log the error
		log.WithFields(logrus.Fields{
			"microservice": "API Gateway",
			"endpoint":     ctx.FullPath(),
			"error":        err.Error(),
		}).Error("Error processing request")

		return
	}

	var form pb.FormRequest

	tempid, _ := strconv.Atoi(empID.(string))

	form.Id = int32(tempid)
	copier.Copy(&form, &tempForm)

	res, _ := c.PostForm(context.Background(), &form)

	if res.Statuscode >= 400 {
		resp := response.ErrorResponse(int(res.Statuscode), res.Message, "", nil)
		ctx.JSON(resp.StatusCode, resp)

		// Log the error
		log.WithFields(logrus.Fields{
			"microservice": "API Gateway",
			"endpoint":     ctx.FullPath(),
			"error":        res.Message,
		}).Error("Error processing request")

		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Form submitted successfully pending for verification", nil)
	ctx.JSON(resp.StatusCode, resp)

	// Log the success
	log.WithFields(logrus.Fields{
		"microservice": "API Gateway",
		"endpoint":     ctx.FullPath(),
		"status_code":  res.Statuscode,
	}).Info("Request processed successfully")
}
