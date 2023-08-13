package routes

import (
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/config"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/form/handler"
	"github.com/ashiqsabith123/shiftsync-grpc-api-gateway/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterFormAdminRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: form.InitClient(c),
	}
	routes := r.Group("/application")
	routes.Use(middleware.AuthenticateAdmin)
	routes.GET("/", svc.GetApplications)
	routes.POST("/approve", svc.ApproveApplication)
	routes.PATCH("/correction", svc.ApplicationCorrection)
}

func (svc *ServiceClient) GetApplications(ctx *gin.Context) {
	handler.ViewApplications(ctx, svc.Client)
}

func (svc *ServiceClient) ApproveApplication(ctx *gin.Context) {
	handler.ApproveApplication(ctx, svc.Client)
}

func (svc *ServiceClient) ApplicationCorrection(ctx *gin.Context) {
	handler.ApplicationCorrection(ctx, svc.Client)
}
