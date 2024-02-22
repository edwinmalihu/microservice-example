package main

import (
	"category-service/model"
	"category-service/route"
)

func main() {

	db, _ := model.DBConnection()
	route.SetupRoutes(db)

}
