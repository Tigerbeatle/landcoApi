package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"encoding/json"
	"github.com/tigerbeatle/landcoApi/models"
	"fmt"
	"log"
	"github.com/gorilla/schema"
	"github.com/imdario/mergo"
)

type ScoopContext struct {
	Db *mongo.Database
}


func (c *ScoopContext) Region(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var regionData models.Region
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&regionData, r.PostForm)
	if err != nil {
		log.Println(err)
	}

	basic := models.BasicJSONReturn{"LandcoAPI", "200", "Region"}
	repo := models.RegionRepo{c.Db.Collection("regions")}
	if(repo.Exists(regionData)){ //replace

		dst := repo.Get(regionData.RegionName)
		err = mergo.Merge(&dst, regionData, mergo.WithOverride)
		if err != nil {
			log.Println(err)
		}
		updateResult := repo.Replace(regionData)
		if(updateResult.MatchedCount == 0){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}else{ //insert
		insertResult := repo.Insert(regionData)
		if(insertResult.InsertedID == ""){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}

func (c *ScoopContext) Parcel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var parcelData models.Parcel
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&parcelData, r.PostForm)
	if err != nil {
		log.Println(err)
	}

	basic := models.BasicJSONReturn{"Ping", "200", "Parcel"}
	repo := models.ParcelRepo{c.Db.Collection("parcels")}

	if(repo.Exists(parcelData)){ //replace

		dst := repo.Get(parcelData.UUID)
		err = mergo.Merge(&dst, parcelData, mergo.WithOverride)
		if err != nil {
			log.Println(err)
		}
		updateResult := repo.Replace(parcelData)
		if(updateResult.MatchedCount == 0){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}else{ //insert
		insertResult := repo.Insert(parcelData)
		if(insertResult.InsertedID == ""){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}

func (c *ScoopContext) GetRegionsByEstate(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["estateid"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}
	estateid := keys[0]

	log.Println("Url Param 'estateid' is: " + string(estateid))

	repo := models.RegionRepo{c.Db.Collection("regions")}
	results := repo.GetByEstateID(estateid)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(results)
}
