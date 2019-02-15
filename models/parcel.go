package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type ParcelRepo struct {
	Coll *mongo.Collection
}

type (
	Parcel struct {
		AccountOwner			Person				`json:"accountOwner"`
		EstateID				string				`json:"estateID"`
		Owner					Person				`json:"owner"`
		Pos						string				`json:"pos"`
		UUID					string				`json:"uuid"` // server assigned
		Name					string				`json:"name"` // Sourced by sl obj llGetParcelDetails
		Desc					string				`json:"desc"` // Sourced by sl obj llGetParcelDetails
		GroupUUID				string				`json:"group"` // Sourced by sl obj (group key) llGetParcelDetails
		Area					int					`json:"area"` // Sourced by sl obj llGetParcelDetails
		SeeAvatars				int					`json:"seeAvatars"` // Sourced by sl obj llGetParcelDetails
		Tenant					Person				`json:"tenant"` // Sourced by sl obj (rental box)
		Prices					[]Price				`json:"prices"` // Sourced by website
		PrimCounts				PrimCounts			`json:"primCount"` // Sourced by sl obj llGetParcelPrimCount
		Flags					ParcelFlags			`json:"flags"` // Sourced by sl obj llGetRegionFlags
		Surl    				string  			`json:"surl"` // Sourced by sl obj (rental box)
		Url     				string  			`json:"url"` // Sourced by sl obj (rental box)
		RentalDate 				time.Time			`json:"rentalDate"`
		RentalDuration			int					`json:"rentalDuration"`
	}

	PrimCounts struct {
		MaxPrims				int					`json:"maxPrims"`
		Total					int					`json:"total"`
		Owner					int					`json:"owner"`
		Group					int					`json:"group"`
		Other					int					`json:"other"`
		Temp					int					`json:"temp"`
	}

	ParcelFlags struct {
		AllowFly					string		`json:"allowFly"`
		AllowScripts				string		`json:"allowScripts"`
		AllowLandmarks				string		`json:"allowLandmarks"`
		AllowTerraform				string		`json:"allowTerraform"`
		AllowDamage					string		`json:"allowDamage"`
		AllowCreateObject			string		`json:"allowCreateObject"`
		UseAccessGroup				string		`json:"useAccessGroup"`
		UseAccessList				string		`json:"useAccessList"`
		UseBanList					string		`json:"useBanList"`
		UseLandPassList				string		`json:"useLandPassList"`
		LocalSoundOnly				string		`json:"localSoundOnly"`
		RestrictPushObject			string		`json:"restrictPushObject"`
		AllowGroupScripts			string		`json:"allowGroupScripts"`
		AllowCreateGroupObjects		string		`json:"allowCreateGroupObjects"`
		AllowAllObjectEntry			string		`json:"allowAllObjectEntry"`
		AllowGroupObjectEntry		string		`json:"allowGroupObjectEntry"`
	}


)
