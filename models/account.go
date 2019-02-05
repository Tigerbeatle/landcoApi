package models

import "github.com/mongodb/mongo-go-driver/mongo"

type AccountRepo struct {
	Coll *mongo.Collection
}

type (
	AccountHolder struct {
		AccountOwner		Person `json:"person" bson:"person"`
		ShareHolders		[]ShareHolder  `json:"shareHolders" bson:"shareHolders"`
		OrganizationName	string `json:"organizationName" bson:"organizationName"`
		AccountUUID			string `json:"accountUUID" bson:"accountUUID"`
		Email				string `valid:"email" json:"email" bson:"email"`
	}

	ShareHolder struct {
		Percentage int	`json:"percentage"  bson:"percentage"`
		ShareHolderName Person `json:"shareHolderName"  bson:"shareHolderName"`
	}
)

