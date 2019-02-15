package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"encoding/json"
	"github.com/tigerbeatle/landcoApi/models"
	"fmt"
	"log"
	"github.com/gorilla/schema"
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

	fmt.Println("regionData.AccountOwner.Name:", regionData.AccountOwner.Name)
	fmt.Println("regionData.EstateName:", regionData.EstateName)
	fmt.Println("regionData.Flags.AllowDamage:", regionData.Flags.AllowDamage)




	basic := models.BasicJSONReturn{"Ping", "200", "Region"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}

func (c *ScoopContext) Parcel(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var dnsEntry models.DnsEntry
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&dnsEntry, r.PostForm)
	if err != nil {
		log.Println(err)
	}

	basic := models.BasicJSONReturn{"Ping", "200", "Parcel"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}