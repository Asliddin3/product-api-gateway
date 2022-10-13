package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Asliddin3/product-api-gateway/config"
	pb "github.com/Asliddin3/product-api-gateway/genproto/product"
	spb "github.com/Asliddin3/product-api-gateway/genproto/store"
)

type IServiceManager interface {
	ProductService() pb.ProductServiceClient
	StoreService() spb.StoreserviceClient
}

type serviceManager struct {
	productService pb.ProductServiceClient
	storeService   spb.StoreserviceClient
}

func (s *serviceManager) StoreService() spb.StoreserviceClient {
	return s.storeService
}

func (s *serviceManager) ProductService() pb.ProductServiceClient {
	return s.productService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connProduct, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.ProductServiceHost, conf.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connStore, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.StoreServiceHost, conf.StoreServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		productService: pb.NewProductServiceClient(connProduct),
		storeService:   spb.NewStoreserviceClient(connStore),
	}

	return serviceManager, nil
}
