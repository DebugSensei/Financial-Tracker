package main

import (
	"financial_tracker/api"
	"financial_tracker/db"
	"log"
)

func main() {
	db.InitDB()
	r := api.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
