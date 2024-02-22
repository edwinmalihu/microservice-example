package main

import (
	"product-service/model"
	"product-service/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)

}
