package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"fmt"
)

type AccountContext struct {
	Db *mongo.Database
}

func (c *AccountContext) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In PING SO mabye before this?")
	basic := models.BasicJSONReturn{"LandcoAPI", "200", "Pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(basic)

}

