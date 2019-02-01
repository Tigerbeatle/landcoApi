package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"fmt"
	"log"
	"github.com/gorilla/schema"
)

type AccountContext struct {
	Db *mongo.Database
}

func (c *AccountContext) Ping(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"LandcoAPI", "200", "Pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}




func (c *AccountContext) DnsRegister(w http.ResponseWriter, r *http.Request){
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

	dnsEntry.SerialNumber = r.Header.Get("X-SecondLife-Object-Key")
	dnsEntry.Owner.Name = r.Header.Get("X-SecondLife-Owner-Name")
	dnsEntry.Owner.UUID = r.Header.Get("X-SecondLife-Owner-Key")
	dnsEntry.Region = r.Header.Get("X-SecondLife-Region")


	basic := models.BasicJSONReturn{"LandcoAPI", "200", "DNS-Registered"}

	repo := models.DnsRepo{c.Db.Collection("dns")}
	if(repo.Exists(dnsEntry)){ //update
		updateResult := repo.Update(dnsEntry)
		if(updateResult.MatchedCount == 0){
			basic = models.BasicJSONReturn{"LandcoAPI", "500", "ErrInternalServer"}
		}
		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	}else{ //insert
		insertResult := repo.Insert(dnsEntry)
		if(insertResult.InsertedID == ""){
			basic = models.BasicJSONReturn{"LandcoAPI", "500", "ErrInternalServer"}
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)



		// Send createDefault Box request
		boxRepo := models.BoxRepo{c.Db.Collection("box")}
		insertResult = boxRepo.CreateDefault(dnsEntry)
		if(insertResult.InsertedID == ""){
			basic = models.BasicJSONReturn{"LandcoAPI", "500", "ErrInternalServer"}
		}
		fmt.Println("Inserted a single box document: ", insertResult.InsertedID)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)
}


