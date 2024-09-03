package main

import (
    "log"    
    "backend/routes"    	
    "backend/utils"
)

func main() {
    // Initialize database	
    DB := utils.InitDB()
	

     // Set up Gin router    
    router := routes.setupRouter(DB)    

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
