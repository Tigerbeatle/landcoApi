package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"fmt"
)


type ControlContext struct {
	Db mongo.Database
}

type DnsRepo struct {
	Coll *mongo.Collection
}



type (

	Parcel struct {
		Surl    string  `json:"surl"    bson:"surl"`
		Url     string  `json:"url"     bson:"url"`
		Name    string  `json:"name"    bson:"name"`
	}

	Person struct {
		UUID    string  `json:"uuid"    bson:"uuid"`
		Name    string  `json:"name"    bson:"name"`
	}

	DnsEntry struct {
		ID                  primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
		SerialNumber        string  `json:"serialNumber"    bson:"serialNumber"`
		Language            string  `json:"language"        bson:"language"`
		Type                string  `json:"type"            bson:"type"`
		Version             float64 `json:"version"         bson:"version"`
		AliveTestCount      int     `json:"aliveTestCount"  bson:"aliveTestCount"`
		RemoveTarget        bool    `json:"removeTarget"    bson:"removeTarget"`
		Blocked             bool    `json:"blocked"         bson:"blocked"`
		AliveTestStatus     string  `json:"aliveTestStatus" bson:"aliveTestStatus"`
		Owner               Person  `json:"owner"           bson:"owner"`
		Parcel              Parcel  `json:"parcel"          bson:"parcel"`
		Region              Parcel  `json:"region"          bson:"regio"`
	}


)


func (r *DnsRepo) Exists(e DnsEntry) bool {
	// look for record via serial number (uuid of rental box)
	var result DnsEntry
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return true
}

func (r *DnsRepo) Insert(e DnsEntry)  *mongo.InsertOneResult{
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}


func (r *DnsRepo) Update(e DnsEntry)  *mongo.UpdateResult{
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	update := bson.D{
		{"$set", bson.D{
			{"aliveTestStatus", e.AliveTestStatus},
			{"aliveTestCount", e.AliveTestCount},
			{"parcel.parcelSurl", e.Parcel.Surl},
			{"parcel.parcelUrl", e.Parcel.Url},
			{"parcel.parcelName", e.Parcel.Name},
		}},
	}

	updateResult, err := r.Coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return updateResult
}

func (r *DnsRepo) Get(e DnsEntry)  DnsEntry{  // NOTE: UNTESTED
	var result DnsEntry
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}
func (r *DnsRepo) Delete(e DnsEntry)  { // NOTE: UNTESTED
	filter := bson.D{{"serialNumber", e.SerialNumber}}
	deleteResults, err := r.Coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Deleted %v document in the collection\n",deleteResults.DeletedCount)
}



