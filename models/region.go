package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
	"context"
)


type (
	Region struct {
		AccountOwner			Person				`json:"accountOwner"`
		EstateID				string				`json:"estateID"`
		EstateName				string				`json:"estateName"`
		SimulatorHostname		string				`json:"simulatorHostname"`
		Name					string				`json:"name"`
		Rating					string				`json:"rating"`
		Status					string				`json:"status"`
		Pos						string				`json:"pos"`
		TimeDilation			string				`json:"timeDilation"`
		FPS						string				`json:"FPS"`
		AgentLimit				string				`json:"agentLimit"`
		DynamicPathfinding		string				`json:"dynamicPathfinding"`
		FrameNumber				string				`json:"frameNumber"`
		CPURatio				string				`json:"CPURatio"`
		Idle					string				`json:"idle"`
		ProductName				string				`json:"productName"`
		ProjectSku				string				`json:"projectSku"`
		StartTime				string				`json:"startTime"`
		SimVersion				string				`json:"simVersion"`
		MaxPrims				string				`json:"maxPrims"`
		ObjectBonus				string				`json:"objectBonus"`
		Flags					RegionFlags			`json:"flags"`
	}

	RegionFlags struct {
		Sandbox					string				`json:"sandbox"`
		AllowDamage				string				`json:"allowDamage"`
		FixedSun				string				`json:"fixedSun"`
		BlockTerraform			string				`json:"blockTerraform"`
		DisableCollision		string				`json:"disableCollision"`
		DisablePhysics			string				`json:"disablePhysics"`
		BlockFly				string				`json:"blockFly"`
		AllowDirectTeleport		string				`json:"allowDirectTeleport"`
		RestrictPushObject		string				`json:"restrictPushObject"`
	}

)


type RegionRepo struct {
	Coll *mongo.Collection
}


func (r *RegionRepo) Exists(e Region) bool {
	// look for record via serial number (uuid of rental box)
	var result Box
	filter := bson.D{{"name", e.Name}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}


func (r *RegionRepo) Insert(e Region)  *mongo.InsertOneResult{
	insertResult, err := r.Coll.InsertOne(context.TODO(), e)
	if err != nil {
		log.Println(err)
	}
	return insertResult
}

func (r *RegionRepo) Get(name string)  Region{
	var result Region
	filter := bson.D{{"name", name}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func (r *RegionRepo) Replace(e Region)  *mongo.UpdateResult{
	filter := bson.D{{"name", e.Name}}
	update := e

	replaceResult, err := r.Coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return replaceResult
}