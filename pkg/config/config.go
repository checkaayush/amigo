package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Load returns Configuration struct
func Load(appName string) (*Configuration, error) {
	var cfg Configuration
	if err := envconfig.Process(appName, &cfg); err != nil {
		return nil, fmt.Errorf("error loading configuration for %s: %s", appName, err)
	}

	return &cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server
	// DB     *Database    `yaml:"database,omitempty"`
	// App    *Application `yaml:"application,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port string `envconfig:"port" default:"5000"`
	// Debug bool   `envconfig:"debug,omitempty"`
}

// // Database holds data necessery for database configuration
// type Database struct {
// 	PSN        string `yaml:"psn,omitempty"`
// }

// // Application holds application configuration details
// type Application struct {
// 	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
// 	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
// }
