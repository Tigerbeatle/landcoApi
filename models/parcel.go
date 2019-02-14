package models

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type ParcelRepo struct {
	Coll *mongo.Collection
}

type (
	Parcel struct {
		UUID				string		`json:"uuid"` // server assigned
		Name				string		`json:"name"` // Sourced by sl obj llGetParcelDetails
		Desc				string		`json:"desc"` // Sourced by sl obj llGetParcelDetails
		Group				string		`json:"group"` // Sourced by sl obj (group key) llGetParcelDetails
		Area				int			`json:"area"` // Sourced by sl obj llGetParcelDetails
		SeeAvatars			int			`json:"seeAvatars"` // Sourced by sl obj llGetParcelDetails
		Tenant				Person		`json:"tenant"` // Sourced by sl obj (rental box)
		Prices				[]Price		`json:"prices"` // Sourced by website
		RadioURL			string		`json:"radioURL"` // Sourced by sl obj llGetParcelMusicURL
		PrimCount			PrimCount	`json:"primCount"` // Sourced by sl obj llGetParcelPrimCount
		PrimCountSimWide	PrimCount	`json:"primCountSimWide"` // Sourced by sl obj llGetParcelPrimCount
		MaxPrims			int			`json:"maxPrims"` // Sourced by sl obj llGetParcelMaxPrims
		MaxPrimsSimWide		int			`json:"maxPrimsSimWide"` // Sourced by sl obj llGetParcelMaxPrims
		Flags				ParcelFlags	`json:"flags"` // Sourced by sl obj llGetRegionFlags
		Surl    			string  	`json:"surl"` // Sourced by sl obj (rental box)
		Url     			string  	`json:"url"` // Sourced by sl obj (rental box)

	}

	PrimCount struct {
		Total	int		`json:"total"` // Sourced by sl obj llGetParcelPrimCount
		Owner	int		`json:"owner"` // Sourced by sl obj llGetParcelPrimCount
		Group	int		`json:"group"` // Sourced by sl obj llGetParcelPrimCount
		Other	int		`json:"other"` // Sourced by sl obj llGetParcelPrimCount
		Temp	int		`json:"temp"` // Sourced by sl obj llGetParcelPrimCount
	}

	ParcelFlags struct {
		AllowFly				int	`json:"AllowFly"` // Sourced by sl obj llGetRegionFlags
		AllowScripts			int	`json:"AllowScripts"` // Sourced by sl obj llGetRegionFlags
		AllowLandmark			int	`json:"AllowLandmark"` // Sourced by sl obj llGetRegionFlags
		AllowTerrafrom			int	`json:"AllowTerrafrom"` // Sourced by sl obj llGetRegionFlags
		AllowDamage				int	`json:"AllowDamage"` // Sourced by sl obj llGetRegionFlags
		AllowCreateObjects		int	`json:"AllowCreateObjects"` // Sourced by sl obj llGetRegionFlags
		UseAccessGroup			int	`json:"UseAccessGroup"` // Sourced by sl obj llGetRegionFlags
		UseAccessList			int	`json:"UseAccessList"` // Sourced by sl obj llGetRegionFlags
		UseBanList				int	`json:"UseBanList"` // Sourced by sl obj llGetRegionFlags
		UseLandPassList			int	`json:"UseLandPassList"` // Sourced by sl obj llGetRegionFlags
		LocalSoundOnly			int	`json:"LocalSoundOnly"` // Sourced by sl obj llGetRegionFlags
		RestrictPushObject		int	`json:"RestrictPushObject"` // Sourced by sl obj llGetRegionFlags
		AllowGroupObjects		int	`json:"AllowGroupObjects"` // Sourced by sl obj llGetRegionFlags
		AllowCreateGroupObjects	int	`json:"AllowCreateGroupObjects"` // Sourced by sl obj llGetRegionFlags
		AllowAllObjectEntry		int	`json:"AllowAllObjectEntry"` // Sourced by sl obj llGetRegionFlags
		AllowGroupObjectEntry	int	`json:"AllowGroupObjectEntry"` // Sourced by sl obj llGetRegionFlags
	}
)
