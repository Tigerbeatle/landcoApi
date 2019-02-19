package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"github.com/mongodb/mongo-go-driver/bson"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)


type (
	Region struct {
		AccountOwner			Person				`json:"accountOwner"`
		EstateID				string				`json:"estateID"`
		EstateName				string				`json:"estateName"`
		SimulatorHostname		string				`json:"simulatorHostname"`
		RegionName				string				`json:"regionName"`
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
	var result Region
	filter := bson.D{{"regionname", e.RegionName}}
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

func (r *RegionRepo) Get(regionName string)  Region{
	var result Region
	filter := bson.D{{"regionname", regionName}}
	err := r.Coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
	}
	//fmt.Printf("Found a single document: %+v\n", result)
	return result
}

func (r *RegionRepo) Replace(e Region)  *mongo.UpdateResult{
	filter := bson.D{{"regionname", e.RegionName}}
	update := e
	replaceResult, err := r.Coll.ReplaceOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	return replaceResult
}


func (r *RegionRepo) GetByEstateID(estateId string)  []*Region{

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*Region
	filter := bson.D{{"estateid", estateId}}
	cur, err := r.Coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Region
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	//fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results
}


