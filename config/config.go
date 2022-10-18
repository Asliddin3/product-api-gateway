package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment        string
	ProductServiceHost string
	ProductServicePort int
	StoreServiceHost   string
	StoreServicePort   int
	CtxTimeout         int
	LogLevel           string
	HTTPPort           string
}

func Load() Config {
	c := Config{}
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8070"))
	c.ProductServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "localhost"))
	c.ProductServicePort = cast.ToInt(getOrReturnDefault("PRODUCT_SERVICE_PORT", 9000))
	c.StoreServiceHost = cast.ToString(getOrReturnDefault("STORE_SERVICE_HOST", "localhost"))
	c.StoreServicePort = cast.ToInt(getOrReturnDefault("STORE_SERVICE_PORT", 8000))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c

}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
