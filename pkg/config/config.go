package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Configuration holds data necessery for configuring application
type Configuration struct {
	DB     *Database `envconfig:"mongodb"`
	Server *Server   `envconfig:"server"`
}

// Server holds data necessary for server configuration
type Server struct {
	Port string `envconfig:"port"`
}

// Database holds data necessary for database configuration
type Database struct {
	Host     string `envconfig:"host" default:"mongodb:27017"`
	Name     string `envconfig:"name" default:"amigo"`
	Username string `envconfig:"username" default:""`
	Password string `envconfig:"password" default:""`
	Timeout  int    `envconfig:"timeout" default:"10"`
}

// Load returns Configuration struct
func Load(appName string) (*Configuration, error) {
	var cfg Configuration
	if err := envconfig.Process(appName, &cfg); err != nil {
		return nil, fmt.Errorf("error loading configuration for %s: %s", appName, err)
	}

	return &cfg, nil
}
