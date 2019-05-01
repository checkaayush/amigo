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
	ServerPort string `split_words:"true" default:"5000"`

	MongodbName     string `split_words:"true" default:"amigo"`
	MongodbHost     string `split_words:"true" default:"mongodb"`
	MongodbUsername string `split_words:"true"`
	MongodbPassword string `split_words:"true"`
	MongodbTimeout  int    `split_words:"true" default:"60"`

	// Server *Server
	// DB     *Database
	// App    *Application `yaml:"application,omitempty"`
}

// Server holds data necessery for server configuration
// type Server struct {
// 	Port string `envconfig:"port" default:"5000"`
// }

// // Database holds data necessery for database configuration
// type Database struct {
// 	Host     string `envconfig:"host" default:"mongodb"`
// 	Name     string `envconfig:"dbname" default:"amigo"`
// 	Username string `envconfig:"username,omitempty"`
// 	Password string `envconfig:"password,omitempty"`
// }

// // Application holds application configuration details
// type Application struct {
// 	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
// 	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
// }
