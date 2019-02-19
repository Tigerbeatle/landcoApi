package main

import (
	"github.com/tigerbeatle/landcoApi/models"
	"github.com/justinas/alice"
	"github.com/tigerbeatle/landcoApi/middleware"
	"github.com/tigerbeatle/landcoApi/routes"
	controller "github.com/tigerbeatle/landcoApi/controllers"
	"net/http"
	"log"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var s models.Specification
	err := envconfig.Process("landco", &s)
	if err != nil {
		log.Fatal(err.Error())
	}


	db := models.NewMongoDB()

	// Lets set some routes

	commonHandlers := alice.New(middleware.RecoverHandler, middleware.AcceptHandler)
	router := routes.NewRouter()


	appA := controller.AccountContext{db.Database}
	appH := controller.HomeContext{db.Database}
	appB := controller.BoxContext{db.Database}
	appS := controller.ScoopContext{db.Database}
	//appI := controller.InvPublicContext{db.Database}
	//appIPi := controller.InvPrivateContext{db.Database}

	//
	// PUBLIC
	//

	// root
	router.Get("/", commonHandlers.ThenFunc(appH.HomeHandler))
	router.Get("/api/1.0/acct/ping", commonHandlers.ThenFunc(appA.Ping))
	router.Post("/api/1.0/dns/register", commonHandlers.ThenFunc(appA.DnsRegister))

	router.Post("/api/1.0/box/record", commonHandlers.ThenFunc(appB.Record))

	router.Post("/api/1.0/box/updateBox", commonHandlers.ThenFunc(appB.UpdateBox))


	router.Post("/api/1.0/scoop/region", commonHandlers.ThenFunc(appS.Region))
	router.Post("/api/1.0/scoop/parcel", commonHandlers.ThenFunc(appS.Parcel))

	router.Get("/api/1.0/scoop/region/getRegionsByEstate", commonHandlers.ThenFunc(appS.GetRegionsByEstate))




	// account (accessable by anyone)
	//router.Post("/api/1.0/acct/createUserProfile", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.User{})).ThenFunc(appA.CreateUserProfile))

	// Public Inventory operations
	//router.Post("/api/1.0/inv/public/get", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appI.Get))
	//router.Post("/api/1.0/inv/public/getByUser", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appI.GetByUser))



	//
	// PRIVATE
	//



	// inventory Private Operations
	// todo add middleware to test valid login
	//router.Post("/api/1.0/u/inv/add", commonHandlers.Append(middleware.ContentTypeHandler, middleware.AuthorizationHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appIPi.Add))
	//router.Post("/api/1.0/u/inv/remove", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appIPi.Remove))

	//router.Post("/api/1.0/u/inv/get", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appIPi.Get))
	//router.Post("/api/1.0/u/inv/getByUser", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appIPi.GetByUser))


	//router.Post("/api/1.0/u/inv/update", commonHandlers.Append(middleware.ContentTypeHandler, middleware.BodyHandler(models.Item{})).ThenFunc(appIPi.Update))



	log.Println("API Starting on Port:",s.Port,"...")


	log.Fatal(http.ListenAndServe(":8002", router))

	log.Println("API Stopped")


}





