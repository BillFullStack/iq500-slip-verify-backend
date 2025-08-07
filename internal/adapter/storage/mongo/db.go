package mongo

import (
	"context"
	"main/internal/adapter/config"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CONNECT_TIME_OUT = 30
)

type Resource struct {
	DB *mongo.Database
	// MessageDB *mongo.Database
}

type UpdateOperation struct {
	Filter bson.M
	Update bson.M
}

func New(config *config.DB) (*Resource, error) {
	_ = godotenv.Load()

	client, err := connectDB(config)
	if err != nil {
		return nil, err
	}

	// messageClient, err := connectDB(messageConfig)
	// if err != nil {
	// 	return nil, err
	// }

	color.Green("Connect database successfully")

	return &Resource{
		DB: client.Database(config.Name),
	}, nil
}

func connectDB(config *config.DB) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	ctx, cancel := context.WithTimeout(context.Background(), CONNECT_TIME_OUT*time.Second)
	defer cancel()

	client, err = mongo.Connect(
		ctx,
		options.Client().ApplyURI(config.Connection),
		options.Client().SetMinPoolSize(1),
		options.Client().SetMaxPoolSize(2),
	)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (r *Resource) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), CONNECT_TIME_OUT*time.Second)
	defer cancel()

	if err := r.DB.Client().Disconnect(ctx); err != nil {
		color.Red("Close connection 'MainDB' falure, Something wrong...")
	}

	color.Cyan("Close connection successfully")
}
