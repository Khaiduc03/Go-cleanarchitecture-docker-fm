package configuration

import (
	"FM/src/core/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type ConfigImpl struct{}

func (config *ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func NewConfig(filename ...string) Config {
	err := godotenv.Load(filename...)
	exception.PanicLogging(err)
	return &ConfigImpl{}
}
