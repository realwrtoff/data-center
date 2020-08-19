package router

import (
	"github.com/gin-gonic/gin"
	"github.com/realwrtoff/data-center/internal/service"
)

func InitCompanyRouter(Router *gin.Engine, svc *service.Service) {
	UserRouter := Router.Group("company")
	// UserRouter.Use(middleware.JWT())
	{
		UserRouter.GET("search", svc.CompanySearch)
		UserRouter.GET("detail", svc.CompanyDetail)
		UserRouter.GET("person", svc.PersonSearch)
	}
}
