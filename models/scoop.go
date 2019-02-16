package models



import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
	"context"
)


type ScoopRepo struct {
	Coll *mongo.Collection
}

func (r *ScoopRepo) RegionExists(e Region) bool {
	// look for record via serial number (uuid of rental box)
	var result Box
	filter := bson.D{{"name", e.Name}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}


func (r *ScoopRepo) RegionInsert(e Region)  *mongo.InsertOneResult{
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}

func (r *ScoopRepo) RegionGet(name string)  Region{
	var result Region
	filter := bson.D{{"name", name}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func (r *ScoopRepo) RegionReplace(e Region)  *mongo.UpdateResult{
	filter := bson.D{{"name", e.Name}}
	update := e

	replaceResult, err := r.Coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return replaceResult
}