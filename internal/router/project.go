package router

import (
	"github.com/gin-gonic/gin"
	"github.com/realwrtoff/data-center/internal/service"
)

func InitProjectRouter(Router *gin.Engine, svc *service.Service) {
	ProjectRouter := Router.Group("project")
	NewSeedGroup := ProjectRouter.Group("newseed")
	{
		NewSeedGroup.GET("search", svc.NewSeedSearch)
		NewSeedGroup.GET("detail", svc.NewSeedDetail)
		NewSeedGroup.GET("invest", svc.NewSeedInvest)
	}
	AiHeHuoGroup := ProjectRouter.Group("aihehuo")
	{
		AiHeHuoGroup.GET("search", svc.AiHeHuoSearch)
		AiHeHuoGroup.GET("detail", svc.AiHeHuoDetail)
		AiHeHuoGroup.GET("publisher", svc.AiHeHuoPublisher)
	}
}
