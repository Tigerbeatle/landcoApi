package controllers
import (
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"encoding/json"
	"github.com/mongodb/mongo-go-driver/mongo"
)


type HomeContext struct {
	Db *mongo.Database

}

func (c *HomeContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"LandcoAPI", "Home", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}

func (c *HomeContext) LoginHandler(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"LandcoAPI", "Login", ""}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)
}