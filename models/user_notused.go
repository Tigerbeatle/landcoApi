package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	//"github.com/mongodb/mongo-go-driver/mongo/options"
	//"time"
	//"errors"
	"log"
	"context"
)

type (

	User struct {
		ID         primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
		UUID     string `json:"uuid"      bson:"uuid"`
	}

)

type UserRepo struct {
	Coll *mongo.Collection
}

type UserResource struct {
	Data User `json:"data"`
}


type UsersCollection struct {
	Data []User `json:"data"`
}


func (r *UserRepo) UserExist(user *User) (bool) {
	var result User
	filter := bson.D{{"uuid", user.UUID}}

	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
/*
	cursor, err := r.Coll.Find(
		context.Background(),

		bson.NewDocument(bson.EC.String("uuid", user.UUID)),
	)

	if err != nil {
		log.Fatal(err)
	}


	defer cursor.Close(context.Background())
	itemRead := User{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&itemRead)
		if err != nil {
			log.Fatal(err)
		}
		return true
	}
	// no user found, return empty string
	return false
*/
}

/*

func (r *UserRepo) Create(user *User) (primitive.ObjectID, error) {
	fmt.Println("------user.UUID:", user.UUID)
	res, err := r.Coll.InsertOne(context.Background(), bson.NewDocument(
		bson.EC.String("uuid", user.UUID),
	))
	return res.InsertedID.(primitive.ObjectID), err
}

*/

/*
func (r *UserRepo) Login(email string, password string) (UserResource, error) {
	result := UserResource{}

	fmt.Println("	Inside Login A")
	fmt.Println("Email:", email)
	cursor, err := r.Coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("passord", password)),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("	Inside Login B")
	defer cursor.Close(context.Background())
	//result := User{}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Data)
		fmt.Println("	Inside Login C")
		fmt.Println("stored Password:", result.Data.Password)
		fmt.Println("stored eamil   :", email)
		fmt.Println("passed Password:", password)
		err = bcrypt.CompareHashAndPassword([]byte(result.Data.Password), []byte(password))
		if err != nil {
			fmt.Println("Found the error")
			return result, err // return err if password doesn't match hashed password stored in db
		}
	}
	fmt.Println("result:", result.Data)
	return result, nil
}
*/