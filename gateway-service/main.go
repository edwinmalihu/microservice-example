package main

import (
	"auth-services/model"
	"auth-services/route"

	_ "auth-services/docs"
)

func main() {
	// pengajuan := &model.SignatureRequest{
	// 	Host:    "api.tapera.digi46.id",
	// 	Path:    "/api/terbit/release",
	// 	Payload: "{\"no_pengajuan\":\"KKP1111220016\",\"no_kartu\":\"123456789\",\"corporate_number\":\"PT SGR\",\"created_by\":\"superadmin-kkp\"}",
	// }
	// // Payload: "{\"kode_satker\":\"007\",\"created_by\":\"superadmin-kkp\",\"nama_lengkap\":\"mamans\",\"limit_card\":\"1000000\"}",

	// json_request, err := json.Marshal(pengajuan)
	// if err != nil {
	// 	log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// }

	// claims := jwt.MapClaims{
	// 	"exp":  time.Now().Add(time.Hour * 3).Unix(),
	// 	"iat":  time.Now().Unix(),
	// 	"data": string(json_request),
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// t, _ := token.SignedString([]byte("{*aKV|~ACD$Gm-Sk"))
	// fmt.Printf("%s\n", t)

	// claims := jwt.MapClaims{
	// 	"exp":      time.Now().Add(time.Hour * 3).Unix(),
	// 	"iat":      time.Now().Unix(),
	// 	"username": "whs_admin1",
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// t, _ := token.SignedString([]byte("{*aKV|~ACD$Gm-Sk"))
	// fmt.Printf("%s\n", t)

	// @title	Synapsis E-commerce API
	// @description	Service API in Go Using Gin framework
	// @host	43.218.113.72:8081

	// @port	running in port 8081
	// @BasePath	/api

	db, _ := model.DBConnection()
	route.SetupRoutes(db)

}
