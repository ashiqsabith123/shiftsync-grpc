package routes

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/middleware"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/handler"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/punch/pb"
	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.PunchServiceClient
}

func RegisterPunchEmployeeRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: punch.InitClient(c),
	}

	route := r.Group("/employee/duty")
	route.Use(middleware.AuthenticateUser)
	route.GET("/", svc.GetDuty)
	route.GET("/punchin", svc.PunchIn)
	route.POST("/punchin", svc.VerifyOtpPunchin)
	route.GET("/punchout", svc.PunchOut)

}

func (svc *ServiceClient) GetDuty(c *gin.Context) {
	handler.GetDuty(c, svc.Client)
}

func (svc *ServiceClient) PunchIn(c *gin.Context) {
	handler.PunchIn(c, svc.Client)
}

func (svc *ServiceClient) VerifyOtpPunchin(c *gin.Context) {
	handler.VerifyOtpPunchin(c, svc.Client)
}

func (svc *ServiceClient) PunchOut(c *gin.Context) {
	handler.PunchOut(c, svc.Client)
}
