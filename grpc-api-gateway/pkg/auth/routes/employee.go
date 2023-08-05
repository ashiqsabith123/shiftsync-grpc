package routes

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/handler"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/auth/pb"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func RegisterAuthEmployeeRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: auth.InitClient(c),
	}
	route := r.Group("/employee")
	route.POST("/signup", middleware.CheckSignupCookie, svc.PostSignup)
	route.POST("/verify-otp", middleware.CheckCookieIsPresent, svc.VerifyOtp)
	route.POST("/signin", svc.PostLogin)
}

func (svc *ServiceClient) PostSignup(ctx *gin.Context) {
	handler.PostSignup(ctx, svc.Client)
}

func (svc *ServiceClient) VerifyOtp(ctx *gin.Context) {
	handler.VerifyOtp(ctx, svc.Client)
}

func (svc *ServiceClient) PostLogin(ctx *gin.Context) {
	handler.PostLogin(ctx, svc.Client)
}
