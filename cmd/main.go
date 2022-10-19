package main

import (
	"github.com/Asliddin3/product-api-gateway/api"
	"github.com/Asliddin3/product-api-gateway/config"
	"github.com/Asliddin3/product-api-gateway/pkg/logger"
	"github.com/Asliddin3/product-api-gateway/services"
	// _ "github.com/swaggo/swag/example/celler/docs"
)

func main() {
	// r := gin.Default()

	// c := controller.NewController()

	// v1 := r.Group("/api/v1")
	// {
	// 	accounts := v1.Group("/accounts")
	// 	{
	// 		accounts.GET(":id", c.ShowAccount)
	// 		accounts.GET("", c.ListAccounts)
	// 		accounts.POST("", c.AddAccount)
	// 		accounts.DELETE(":id", c.DeleteAccount)
	// 		accounts.PATCH(":id", c.UpdateAccount)
	// 		accounts.POST(":id/images", c.UploadAccountImage)
	// 	}
	// 	//...
	// }
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run(":8080")
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
