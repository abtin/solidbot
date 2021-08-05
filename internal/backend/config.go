package backend

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenAddress string `default:"0.0.0.0"`
	ListenPort    int    `default:"8080"`
}

func GetConfig(configPrefix string) (Config, error) {
	var be Config
	err := envconfig.Process(configPrefix, &be)
	if err != nil {
		return Config{}, err
	}
	return be, nil
}

func (b Config) String() string {
	format := "Server Config:\n\tListenAddress: %s\n\tListenPort: %d\n"
	return fmt.Sprintf(format, b.ListenAddress, b.ListenPort)
}
