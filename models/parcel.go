package models

import "github.com/mongodb/mongo-go-driver/mongo"

type ParcelRepo struct {
	Coll *mongo.Collection
}

type (
	Parcel struct {
		AccountUUID string	`json:"accountUUID"`
		Name        string	`json:"name"`
		Desc        string	`json:"desc"`
		Owner       Person	`json:"owner"`
		GroupUUID   string	`json:"groupUUID"`
		Area        int		`json:"area"`
		UUID		string	`json:"uuid"`
		SeeAvatars	int		`json:"seeAvatars"`
		Surl    	string  `json:"surl"    bson:"surl"`
		Url     	string  `json:"url"     bson:"url"`
	}
)
