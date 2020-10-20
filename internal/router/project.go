package router

import (
	"github.com/gin-gonic/gin"
	"github.com/realwrtoff/data-center/internal/service"
)

func InitProjectRouter(Router *gin.Engine, svc *service.Service) {
	ProjectRouter := Router.Group("project")
	NewSeedGroup := ProjectRouter.Group("newseed")
	{
		NewSeedGroup.GET("search", svc.ProjectSearch)
		NewSeedGroup.GET("detail", svc.ProjectDetail)
	}
}
