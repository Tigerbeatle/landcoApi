package models

import (
	"time"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
	"context"
)



type (
	Parcel struct {
		AccountOwner			Person				`json:"accountOwner"`
		EstateID				string				`json:"estateID"`
		RegionName				string				`json:"regionName"`
		Owner					Person				`json:"owner"`
		Pos						string				`json:"pos"`
		UUID					string				`json:"uuid"`
		Name					string				`json:"name"`
		Desc					string				`json:"desc"`
		GroupUUID				string				`json:"group"`
		Area					int					`json:"area"`
		SeeAvatars				int					`json:"seeAvatars"`
		Tenant					Person				`json:"tenant"`
		Prices					[]Price				`json:"prices"`
		PrimCounts				PrimCounts			`json:"primCount"`
		Flags					ParcelFlags			`json:"flags"`
		RentalDate 				time.Time			`json:"rentalDate"`
		RentalDuration			int					`json:"rentalDuration"`
		Surl    string  `json:"surl"    bson:"surl"`
		Url     string  `json:"url"     bson:"url"`
	}

	PrimCounts struct {
		MaxPrims				int					`json:"maxPrims"`
		Total					int					`json:"total"`
		Owner					int					`json:"owner"`
		Group					int					`json:"group"`
		Other					int					`json:"other"`
		Temp					int					`json:"temp"`
	}

	ParcelFlags struct {
		AllowFly					string		`json:"allowFly"`
		AllowScripts				string		`json:"allowScripts"`
		AllowLandmarks				string		`json:"allowLandmarks"`
		AllowTerraform				string		`json:"allowTerraform"`
		AllowDamage					string		`json:"allowDamage"`
		AllowCreateObject			string		`json:"allowCreateObject"`
		UseAccessGroup				string		`json:"useAccessGroup"`
		UseAccessList				string		`json:"useAccessList"`
		UseBanList					string		`json:"useBanList"`
		UseLandPassList				string		`json:"useLandPassList"`
		LocalSoundOnly				string		`json:"localSoundOnly"`
		RestrictPushObject			string		`json:"restrictPushObject"`
		AllowGroupScripts			string		`json:"allowGroupScripts"`
		AllowCreateGroupObjects		string		`json:"allowCreateGroupObjects"`
		AllowAllObjectEntry			string		`json:"allowAllObjectEntry"`
		AllowGroupObjectEntry		string		`json:"allowGroupObjectEntry"`
	}

)

type ParcelRepo struct {
	Coll *mongo.Collection
}


func (r *ParcelRepo) Exists(e Parcel) bool {
	var result Parcel
	filter := bson.D{{"uuid", e.UUID}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}


func (r *ParcelRepo) Insert(e Parcel)  *mongo.InsertOneResult{
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}

func (r *ParcelRepo) Get(UUID string)  Parcel{
	var result Parcel
	filter := bson.D{{"uuid", UUID}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func (r *ParcelRepo) Replace(e Parcel)  *mongo.UpdateResult {
	filter := bson.D{{"uuid", e.UUID}}
	update := e
	replaceResult, err := r.Coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return replaceResult
}