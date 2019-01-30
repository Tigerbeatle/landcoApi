package controllers

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"fmt"
	"log"
	"github.com/tigerbeatle/landcoApi/models"
	"github.com/gorilla/schema"
)


type BoxContext struct {
	Db *mongo.Database
}


func (c *BoxContext) SetPrice(w http.ResponseWriter, r *http.Request){
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

	fmt.Println("print1:",box.Price1)



}