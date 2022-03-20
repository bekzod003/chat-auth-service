package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	mongoHost     string
	mongoPort     string
	mongoDatabase string

	rpcPort string
}

func (c Config) Load() Config {
	c.mongoHost = cast.ToString(getOrReturnDefaultValue("MONGO_HOST", "localhost"))
	c.mongoPort = cast.ToString(getOrReturnDefaultValue("MONGO_PORT", "27017"))
	c.mongoDatabase = cast.ToString(getOrReturnDefaultValue("MONGO_DATABASE", "auth_service"))

	c.rpcPort = cast.ToString(getOrReturnDefaultValue("RPC_PORT", "8001"))
	return c
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

func (c Config) GetMongoHost() string {
	return c.mongoHost
}

func (c Config) GetMongoPort() string {
	return c.mongoPort
}
