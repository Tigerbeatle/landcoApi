package models

import "github.com/mongodb/mongo-go-driver/mongo"

type RegionRepo struct {
	Coll *mongo.Collection
}

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
