package main

import (
	"millieMind/db"
	"millieMind/middleware"
)

func main() {
	r := middleware.SetupRouter()
	db.ConnectDatabase()
	r.Run(":8080")
}
