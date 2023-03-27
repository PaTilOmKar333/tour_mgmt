package main

import (
	"net/http"
	"tour_mgmt/app"
	"tour_mgmt/config"
	"tour_mgmt/handler"
)

func main() {
	// intialise config
	// coonnect with database
	// depedecies initialconfig.Load()
	config.Load()
	app.Init()
	defer app.Close()

	handler.InitDependancies()

	http.ListenAndServe(":8080", handler.Router())

}
