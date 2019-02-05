package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type TransactionRepo struct {
	Coll *mongo.Collection
}

type (
	Transaction struct {
		TransactionUUID string		`json:"transactionUUID"`
		AccountUUID 	string		`json:"accountUUID"`
		Amount			int			`json:"amount"`
		Purpose			string		`json:"purpose"` // ex: rent payment, land purchase, back payment
		SourceObject	string		`json:"sourceObject"` // what object created the transaction
		CreatedAt 		time.Time	`json:"created_at"`
		UpdatedAt 		time.Time	`json:"updated_at"`
		Description 	string		`json:"description"`
		Payee 			Person		`json:"payee"` // The person who made the payment
		ParcelUUID		string		`json:"parcelUUID"`
		ResultUUID		string		`json:"resultUUID"` // transaction result key returned by llTransferLindenDollars
	}
)


