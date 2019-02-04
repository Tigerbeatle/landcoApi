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
	Price struct {
		Amount	 int	 `json:"amount"  bson:"amount"`
		Duration float32 `json:"duration"  bson:"duration"`
	}

	Box struct {
		ProfitShare  string       `json:"profitShare"  bson:"profitShare"`
		Prices       []Price      `json:"prices" bson:"prices"`
		SerialNumber string       `json:"serialNumber"    bson:"serialNumber"`
		ShareOwners  []ShareOwner `json:"shareOwners" bson:"shareOwners"`
	}

	UpdateBoxRequest struct {  // not stored in database
		Type   string
		Status string
		Box      Box
	}

	ShareOwner struct {
		Primary	string  `json:"primary"  bson:"primary"`
		Percentage int	`json:"percentage"  bson:"percentage"`
		UUID    string  `json:"uuid"    bson:"uuid"`
		Name    string  `json:"name"    bson:"name"`

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
	return true
}

func (r *BoxRepo) Insert(e Box)  *mongo.InsertOneResult{
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}

func (r *BoxRepo) Replace(e Box)  *mongo.UpdateResult{
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	update := e

	replaceResult, err := r.Coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return replaceResult
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
	fmt.Printf("Deleted %v document in the collection\n",deleteResults.DeletedCount)
}

func (r *BoxRepo) CreateDefault(e DnsEntry) *mongo.InsertOneResult{
	var box Box
	box.SerialNumber = e.SerialNumber
	box.ProfitShare = "false"

	insertResult, err := r.Coll.InsertOne(context.TODO(), box)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("box insertResult:", insertResult)
	return insertResult
}