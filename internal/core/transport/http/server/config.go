package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig: "SHUTDOWN_TIMEOUT" required:"true"`
}
func newConfig()(Config, error){
	var config Config
	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("process env config: %w", err)
	}
	return config, nil
}
func NewConfigMust() Config{
	config, err := newConfig()
	if err != nil {
		err = fmt.Errorf("Failed to create Config, %w", err)
		panic(err)
	}
	return config
}