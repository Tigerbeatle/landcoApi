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
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	//client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database(s.Database)

	/*
	// Pull in the configuration.
	var mongoConfig connstring.ConnString
	if err := envconfig.Process("mgo", &mongoConfig); err != nil {
		log.Fatalf("Startup Err")
		//return err
	}

	// set up conn to mongoHelper
	client, err := mongo.NewClientFromConnString(mongoConfig)
	if err != nil { log.Fatal(err) }
	// Now connect to the db
	err = client.Connect(context.TODO())
	if err != nil { log.Fatal(err) }

	db := client.Database(mongoConfig.Database)
*/
	return &DB{db}
}