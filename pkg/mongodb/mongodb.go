package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// New creates new connection to MongoDB
func New(host string, dbName string, username string, password string, timeout int) (*mongo.Database, error) {
	cred := options.Credential{
		AuthSource: dbName,
		Username:   username,
		Password:   password,
	}
	clientOptions := options.Client()
	clientOptions.SetAuth(cred)
	clientOptions.SetDirect(true)
	clientOptions.SetHosts([]string{host})
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	Database := client.Database(dbName)
	return Database, nil
}
