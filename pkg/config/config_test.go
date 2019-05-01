package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkaayush/amigo/pkg/config"
)

const AppName = "amigo"

func TestLoad(t *testing.T) {
	cases := []struct {
		name	string
		wantData	*config.Configuration
		wantErr	bool
	}{
		{
			name: "Success",
			wantData: &config.Configuration{
				DB: &config.Database{
					Host: "mongodb:27017",
					Name: "amigo",
					Username: "",
					Password: "",
					Timeout: 10,
				},
				Server: &config.Server{
					Port: "5000",
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			cfg, _ := config.Load(AppName)
			assert.Equal(t, tt.wantData, cfg)
		})
	}
}
