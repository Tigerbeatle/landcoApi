package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
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

	}
)