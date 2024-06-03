package main

import (
	"financial_tracker/api"
	"financial_tracker/db"
	"log"
)

func main() {
	db.ConnectDB()
	defer db.DB.Close()

	r := api.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
