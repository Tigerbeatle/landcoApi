package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	//"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/kelseyhightower/envconfig"
	"log"
	"context"
	"time"
	"fmt"
)

// DB abstraction
type DB struct {
	*mongo.Database
}

func NewMongoDB() *DB {



	var s Specification
	err := envconfig.Process("landco", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Using Database:",s.Database)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(s.Database)

	fmt.Println("Connected to MongoDB!")

	return &DB{db}
}