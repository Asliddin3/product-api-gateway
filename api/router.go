package api

import (
	_ "github.com/Asliddin3/product-api-gateway/api/docs" //swag
	v1 "github.com/Asliddin3/product-api-gateway/api/handlers/v1"
	"github.com/Asliddin3/product-api-gateway/config"
	"github.com/Asliddin3/product-api-gateway/pkg/logger"
	"github.com/Asliddin3/product-api-gateway/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
// @title           Product api
// @version         1.0
// @description     This is shopping server api server
// @termsOfService  not much usefull

// @contact.name   Asliddin
// @contact.url    https://t.me/asliddindeh
// @contact.email  asliddinvstalim@gmail.com

// @host      localhost:8070
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/products/fullinfo", handlerV1.CreateProductFullInfo)
	api.POST("/products", handlerV1.CreateProduct)
	api.GET("/products/:id", handlerV1.GetProduct)
	api.DELETE("product/delete/:id",handlerV1.DeleteProduct)
	api.POST("/products/type", handlerV1.CreateType)
	api.POST("/products/category", handlerV1.CreateType)
	api.GET("/products/info", handlerV1.GetProducts)
	api.PATCH("/products/update", handlerV1.UpdateProduct)
	api.POST("/store", handlerV1.CreateStore)
	api.GET("/store/:id", handlerV1.GetStore)
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
