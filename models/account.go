package models

import "github.com/mongodb/mongo-go-driver/mongo"

type AccountRepo struct {
	Coll *mongo.Collection
}

type (
	AccountHolder struct {
		AccountOwner		Person			`json:"accountOwner"`
		ShareHolders		[]ShareHolder	`json:"shareHolders"`
		OrganizationName	string 			`json:"organizationName"`
		AccountUUID			string 			`json:"accountUUID"`
		Email				string 			`valid:"email" json:"email"`
	}


	ShareHolder struct {
		Percentage 			int				`json:"percentage"`
		ShareHolderName 	Person 			`json:"shareHolderName"`
	}
)

