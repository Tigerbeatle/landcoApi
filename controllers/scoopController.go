package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"encoding/json"
	"github.com/tigerbeatle/landcoApi/models"
	"fmt"
)

type ScoopContext struct {
	Db *mongo.Database
}


func (c *ScoopContext) Region(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	basic := models.BasicJSONReturn{"Ping", "200", "Region"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}

func (c *ScoopContext) Parcel(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	basic := models.BasicJSONReturn{"Ping", "200", "Parcel"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}