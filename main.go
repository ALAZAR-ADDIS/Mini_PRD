package main

import (
	"Mini_PRD/Delivery/Router"
	"Mini_PRD/Repository"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	db := Repository.ConnectDb("admin", "admin")

	Router.InitRoute(db, route)

	route.Run("localhost:8000")
	// Repository.Disconnect(db)
}
