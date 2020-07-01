package main

import (
	"server/db"
	"github.com/gin-gonic/gin"
	"server/router"
	"fmt"
	
)

const port = 3000
func main() {
	db.Connect()
	defer db.Close()
	// Creates a router without any middleware by default
	app := gin.New()
	router.Run(
		app,
	)
	app.Run(":" + fmt.Sprintf("%d", port))
}