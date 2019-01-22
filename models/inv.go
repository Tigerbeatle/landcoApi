package models

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/satori/go.uuid"
	"log"
	"fmt"
)

type (

	Item struct {
		ID         objectid.ObjectID   `json:"id" bson:"_id,omitempty"`
		Name    string `valid:"email" json:"name"       bson:"name" `
		Description string `json:"description"  bson:"description"`
		UUID     string `json:"UUID"      bson:"UUID"`
		UserUUID     string `json:"userUUID"      bson:"userUUID"`
		CollectionUUID     string `json:"collectionUUID"      bson:"collectionUUID"`
		Private      PrivateMetaData	`json:"private" bson:"private"`
	}

	PublicItem struct {
		ID         objectid.ObjectID   `json:"id" bson:"_id,omitempty"`
		Name    string `valid:"email" json:"name"       bson:"name" `
		Description string `json:"description"  bson:"description"`
		UUID     string `json:"UUID"      bson:"UUID"`
		UserUUID     string `json:"userUUID"      bson:"userUUID"`
		CollectionUUID     string `json:"collectionUUID"      bson:"collectionUUID"`
	}

	PrivateMetaData struct {
		Serial  string `json:"serial"      bson:"serial"`
	}
)

type InvRepo struct {
	Coll *mongo.Collection
}

func (r *InvRepo) Create(item *Item) (objectid.ObjectID, error){
	UUID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	res, err := r.Coll.InsertOne(context.Background(),
		bson.NewDocument(
			bson.EC.String("name", item.Name),
			bson.EC.String("description", item.Description),
			bson.EC.String("UUID", UUID.String()),
			bson.EC.String("userUUID", item.UserUUID),
			bson.EC.String("collectionUUID", item.CollectionUUID),
			bson.EC.SubDocumentFromElements("private",
				bson.EC.String("serial", item.Private.Serial),
			),
		),
	)

	return res.InsertedID.(objectid.ObjectID), err
}

func (r *InvRepo) Remove(UUID string) (int64, error){
	res, err := r.Coll.DeleteOne(context.Background(),
		bson.NewDocument(
			bson.EC.String("UUID", UUID)),
	)
fmt.Println("Delete Count:", res.DeletedCount)
	return res.DeletedCount, err
}

func (r *InvRepo) Get(UUID string) (Item, error){
	cursor, err := r.Coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("UUID", UUID)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	item  := Item{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
	}

	return item, nil
}
func (r *InvRepo) PublicGet(UUID string) (PublicItem, error){
	cursor, err := r.Coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("UUID", UUID)),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	item  := PublicItem{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
	}

	return item, nil
}

func (r *InvRepo) Update(body Item) (Item, error){
	item  := Item{}
	res := r.Coll.FindOneAndUpdate(context.Background(),
		bson.NewDocument(bson.EC.String("UUID", body.UUID)),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("name", body.Name),
				bson.EC.String("description", body.Description),
				bson.EC.String("collectionUUID", body.CollectionUUID),
				bson.EC.SubDocumentFromElements("private",
					bson.EC.String("serial", body.Private.Serial),
				),
			),
		),
		nil,
	)

	err := res.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}

	return item, err
}



