package main

import (
    "log"    
    "backend/routes"
    "backend/models"	
)

func main() {
    // Initialize database	
    models.InitDB()
	

     // Set up Gin router    
    router := routes.SetupRoutes()

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
