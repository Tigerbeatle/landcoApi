package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"fmt"
	"log"
	"github.com/tigerbeatle/landcoApi/models"
	"github.com/gorilla/schema"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"github.com/imdario/mergo"
)


type BoxContext struct {
	Db *mongo.Database
}

//
// Record is only to be used by the web interface. Never upstream from the box itself
// Blanks in the source struct will overwrite values in the dst struck and be saved
//
func (c *BoxContext) Record(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var box models.Box
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&box, r.PostForm)
	if err != nil {
		log.Println(err)
	}


	basic := models.BasicJSONReturn{"LandcoAPI", "200", "Box Price Updated"}

	repo := models.BoxRepo{c.Db.Collection("box")}
	if(repo.Exists(box)){ //replace
		dst := repo.Get(box.SerialNumber)
		err = mergo.Merge(&dst, box, mergo.WithOverride)
		if err != nil {
			log.Println(err)
		}
		updateResult := repo.Replace(box)
		if(updateResult.MatchedCount == 0){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}else{ //insert
		insertResult := repo.Insert(box)
		if(insertResult.InsertedID == ""){
			basic = models.BasicJSONReturn{"Record", "500", "ErrInternalServer"}
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}


func (c *BoxContext) UpdateBox(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	var serial = r.FormValue("serialNumber")
	//fmt.Println("serial:",serial)
	boxRepo := models.BoxRepo{c.Db.Collection("box")}
	box := boxRepo.Get(serial)

	//fmt.Println("box:",box)

	// get dns data for the URL to send the updated info too the box
	dnsRepo := models.DnsRepo{c.Db.Collection("dns")}
	dns := dnsRepo.Get(serial)

	jsonStr, err := json.Marshal(box)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(jsonStr))


	//basic := models.BasicJSONReturn{"UpdateBox", "200", string(jsonStr[:len(jsonStr)])}
	//basic := models.BasicJSONReturn{"UpdateBox", "200", jsonStr}

	//s := fmt.Sprintf("%v", basic)
	//req, err := http.NewRequest("POST", dns.Parcel.Url, bytes.NewBuffer([]byte(s)))
	req, err := http.NewRequest("POST", dns.Parcel.Url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("X-Return-Type", "UpdateBox")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()


	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))


	basic := models.BasicJSONReturn{"UpdateBox", "600", "Box Updated"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)
}

