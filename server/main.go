package main

import (
	"server/db"
	"github.com/gin-gonic/gin"
	"server/router"
	
)

const port = "8888"
func main() {
	db.Connect()
	// Creates a router without any middleware by default
	app := gin.New()
	Routes.
	router.Load(
		app
	)

	app.Run(":" + fmt.Sprintf("%d", port))
}