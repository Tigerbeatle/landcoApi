package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"log"
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


	decoder := json.NewDecoder(r.Body)
	var t dnsEntry
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println("****************",t.Parcel)


	tt := json.NewDecoder(r.Body).Decode(dnsEntry)
	fmt.Println("#######---#########-dnsEntry:",dnsEntry)

	q := r.URL.Query()


	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}


	o := r.Header.Get("X-SecondLife-Object-Key")

	fmt.Println("o:",o)
	fmt.Println("r:",r)
	fmt.Println("+++++b:",b)


	fmt.Println("r.Body:",r.Body)
	fmt.Println("json.NewDecoder(r.Body):",json.NewDecoder(r.Body))
	fmt.Println("tt:",tt)
	fmt.Println("dnsEntry:",dnsEntry)

	version, _ := strconv.ParseFloat(q.Get("version"), 32)
	aliveTestCount, _ := strconv.Atoi(q.Get("aliveTestCount"))
	removeTarget, _ := strconv.ParseBool(q.Get("removeTarget"))
	blocked, _ := strconv.ParseBool(q.Get("blocked"))

	//dnsEntry.SerialNumber = q.Get("serialNumber")
	dnsEntry.SerialNumber = r.Header.Get("X-SecondLife-Object-Key")
	dnsEntry.Language = q.Get("language")
	dnsEntry.Version = version
	dnsEntry.AliveTestCount = aliveTestCount
	dnsEntry.RemoveTarget = removeTarget
	dnsEntry.Blocked = blocked
	dnsEntry.AliveTestStatus = q.Get("aliveTestStatus")
	dnsEntry.Owner.Name = r.Header.Get("X-SecondLife-Owner-Name")
	dnsEntry.Owner.UUID = r.Header.Get("X-SecondLife-Owner-Key")
	dnsEntry.Parcel.Surl = q.Get("parcelSurl")
	dnsEntry.Parcel.Url = q.Get("parcelUrl")
	dnsEntry.Parcel.Name = q.Get("parcelName")
	dnsEntry.Region = r.Header.Get("X-SecondLife-Region")

	fmt.Println("------dnsEntry:",dnsEntry)
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


