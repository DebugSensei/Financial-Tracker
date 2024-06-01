package main

import (
	"financial_tracker/api"
)

func main() {
	r := api.SetupRouter()
	r.Run(":8080")
}
