package main

import (
	"order-service/model"
	"order-service/route"
)

func main() {
	db, _ := model.DBConnection()
	route.SetupRoute(db)

}
