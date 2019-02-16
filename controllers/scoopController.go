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

		dst := repo.Get(regionData.Name)
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




	fmt.Println("--regionData.AccountOwner.Name:", regionData.AccountOwner.Name)
	fmt.Println("--regionData.EstateName:", regionData.EstateName)
	fmt.Println("--regionData.Flags.AllowDamage:", regionData.Flags.AllowDamage)




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


	fmt.Println("parcelData.AccountOwner.Name:", parcelData.AccountOwner.Name)
	fmt.Println("parcelData.Owner.Name:", parcelData.Owner.Name)
	fmt.Println("parcelData.Flags.AllowDamage:", parcelData.Flags.AllowDamage)




	basic := models.BasicJSONReturn{"Ping", "200", "Parcel"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}