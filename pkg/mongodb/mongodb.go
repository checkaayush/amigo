package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/checkaayush/amigo/pkg/config"
)

// New creates new connection to MongoDB
func New(cfg *config.Configuration) (*mongo.Database, error) {
	timeout := time.Duration(cfg.DB.Timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Configure database client
	opts := options.Client()
	opts.SetHosts([]string{cfg.DB.Host})
	if cfg.DB.Username != "" && cfg.DB.Password != "" {
		opts.SetAuth(options.Credential{
			AuthSource: cfg.DB.Name,
			Username:   cfg.DB.Username,
			Password:   cfg.DB.Password,
		})
	}
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Blocks for MongoDB server discovery
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.DB.Name)
	return db, nil
}
