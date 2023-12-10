package main

import (
	"fmt"
	"log"
	"os"
	"xyz-multifinance/blueprints"
	"xyz-multifinance/db"
	dbcontroller "xyz-multifinance/db/db-controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode to "release"
	gin.SetMode(gin.ReleaseMode)

	db := db.ConnectGorm()
	serviceController := dbcontroller.Controller(db)
	serviceRouther := blueprints.ServiceRouther(serviceController)

	// Get the port from the environment variable
	port := os.Getenv("SERVER_PORT")
	fmt.Printf("Server is running at port :  %s\n", port)
	serviceRouther.Start(":" + port)
}
