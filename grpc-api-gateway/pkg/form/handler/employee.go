// package handler

// import (
// 	"context"
// 	"fmt"
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

// 	fmt.Println("hello form")

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
	"log"
	"net/http"
	"os"
	"strconv"

	// Replace with the correct import path for your protobuf package

	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/pb"
	models "github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/request"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/models/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func PostForm(ctx *gin.Context, c pb.FormServiceClient) {

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logFile, err := os.OpenFile("pkg/form/logs/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	logger.SetOutput(logFile)

	var tempForm models.Form

	empID, ok := ctx.Get("userId")

	if !ok || empID == "" {
		resp := response.ErrorResponse(500, "Employee id not found in cookie", "", nil)
		ctx.JSON(http.StatusInternalServerError, resp)

		logger.Printf("Service: Form SVC, End Point: %s, Error: %s, Status Code: %d", ctx.FullPath(), resp.Message, resp.StatusCode)

		return
	}

	if err := ctx.ShouldBindJSON(&tempForm); err != nil {
		resp := response.ErrorResponse(400, "Invalid input", err.Error(), tempForm)
		ctx.JSON(400, resp)

		// Log the error

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
		logger.Printf("Service: Form SVC, End Point: %s, Error: %s, Status Code: %d", ctx.FullPath(), resp.Message, resp.StatusCode)

		return
	}

	resp := response.SuccessResponse(int(res.Statuscode), "Form submitted successfully pending for verification", nil)
	ctx.JSON(resp.StatusCode, resp)

	//Log the success

}
