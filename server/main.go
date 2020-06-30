package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/db"
	"server/router"
	
)

const port = "8888"

func main() {
	db.Connect()
	// Creates a router without any middleware by default
	r := gin.New()
	// Routes.
	router.Load(r)
}