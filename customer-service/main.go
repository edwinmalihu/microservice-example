package main

import (
	"customer-service/model"
	"customer-service/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)

}
