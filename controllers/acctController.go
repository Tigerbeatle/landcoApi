package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"github.com/tigerbeatle/landcoApi/models"
	"fmt"
	"github.com/gorilla/context"
	"encoding/json"
)

type AccountContext struct {
	Db *mongo.Database
}

func (c *AccountContext) Ping(w http.ResponseWriter, r *http.Request) {
	basic := models.BasicJSONReturn{"LandcoAPI", "200", "Pong"}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(basic)

}

func (c *AccountContext) CreateUserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Create User r.PostForm :", r.PostForm)
	fmt.Println("Inside Create User context.Get(r, body) :", context.Get(r, "body"))
	body := context.Get(r, "body").(*models.User)

	fmt.Println("+++++body.UUID:",body.UUID)
	rJson := models.BasicJSONReturn{}
	rJson.ReturnType = "registration"
	w.Header().Set("Content-Type", "application/vnd.api+json")

	// 1. Does user exist?
	repo := models.UserRepo{c.Db.Collection("profiles")}
	fmt.Println("body:", body)
	if repo.UserExist(body) {
		// a user already exists
		w.WriteHeader(701)
		rJson.ReturnStatus =  models.ErrUserAlreadyExists.Title
		rJson.Payload = models.ErrUserAlreadyExists.Detail
	}else{
		// no user found. Create user record
		id, err := repo.Create(body)
		if err != nil {
			w.WriteHeader(500)
			rJson.ReturnStatus = models.ErrInternalServer.Title
			rJson.Payload = models.ErrInternalServer.Detail
		}else{
			fmt.Println("objectid.ObjectID ID:", id.String())
			w.WriteHeader(201)
			rJson.ReturnStatus = "success"
			rJson.Payload = id.String()
		}

	}

	json.NewEncoder(w).Encode(rJson)

}


func (c *AccountContext) UserProfile(w http.ResponseWriter, r *http.Request) {
	// verify token  and extract claims before doing anything!
	jwtToken := r.Header.Get("Token")
	claims, ok := models.ExtractClaims(jwtToken)
	if !ok {
		models.WriteError(w, models.ErrUserTokenRejected)
		return
	}
	// Get user's Profile
	repo := models.ProfileRepo{c.Db.Collection("profiles")}
	profile, err := repo.GetPublicProfile(claims["id"].(string))
	if err != nil{
		models.WriteError(w, models.ErrUserNotFound)
	}

fmt.Println("Profile:",profile)

}