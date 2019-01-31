package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
	"context"
	"fmt"
)


type BoxRepo struct {
	Coll *mongo.Collection
}


type (

	Box struct {
		Price1      int     `json:"price1"  bson:"price1"`
		Price2      int     `json:"price2"  bson:"price2"`
		Price3      int     `json:"price3"  bson:"price3"`
		Price4      int     `json:"price4"  bson:"price4"`
		SerialNumber    string  `json:"serialNumber"    bson:"serialNumber"`
	}
)



func (r *BoxRepo) Exists(e Box) bool {
	// look for record via serial number (uuid of rental box)
	var result Box
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return true
}

func (r *BoxRepo) Insert(e Box)  *mongo.InsertOneResult{
	//fmt.Println("box:",e)
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}


func (r *BoxRepo) Update(e Box)  *mongo.UpdateResult{
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	update := bson.D{
		{"$set", bson.D{
			{"price1", e.Price1},
			{"price2", e.Price2},
			{"price3", e.Price3},
			{"price4", e.Price4},
		}},
	}

	updateResult, err := r.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return updateResult
}

func (r *BoxRepo) Get(serialNumber string)  Box{  // NOTE: UNTESTED
	var result Box
	filter := bson.D{{"serialNumber", serialNumber}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func (r *BoxRepo) Delete(serialNumber string)  { // NOTE: UNTESTED
	filter := bson.D{{"serialNumber", serialNumber}}
	deleteResults, err := r.Coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Deleted %v document in the collection\n",deleteResults.DeletedCount)
}