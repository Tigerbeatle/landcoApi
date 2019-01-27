package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"fmt"
	"strconv"
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
	dnsEntry := models.DnsEntry{}
	q := r.URL.Query()

	version, _ := strconv.ParseFloat(q.Get("version"), 32)
	aliveTestCount, _ := strconv.Atoi(q.Get("aliveTestCount"))
	removeTarget, _ := strconv.ParseBool(q.Get("removeTarget"))
	blocked, _ := strconv.ParseBool(q.Get("blocked"))

	dnsEntry.SerialNumber = q.Get("serialNumber")
	dnsEntry.Language = q.Get("language")
	dnsEntry.Version = version
	dnsEntry.AliveTestCount = aliveTestCount
	dnsEntry.RemoveTarget = removeTarget
	dnsEntry.Blocked = blocked
	dnsEntry.AliveTestStatus = q.Get("aliveTestStatus")
	dnsEntry.Owner.Name = q.Get("ownerName")
	dnsEntry.Owner.UUID = q.Get("ownerUUID")
	dnsEntry.Parcel.Surl = q.Get("parcelSurl")
	dnsEntry.Parcel.Url = q.Get("parcelUrl")
	dnsEntry.Parcel.Name = q.Get("parcelName")

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
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)
}


