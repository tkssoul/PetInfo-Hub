package main

import (
    "log"    
    "backend/routes"    	
    "backend/models"
    "github.com/gin-contrib/cors"
)

func main() {
    // Initialize database
    DB := models.InitDB()
	

     // Set up Gin router    
    
    router := routes.SetupRouter(DB)    

    // 使用rs/cors库来配置CORS  
    // 自定义 CORS 配置
    corsConfig := cors.Config{
        AllowOrigins: []string{"*"}, // 允许所有域名        
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
    }

    // 使用自定义 CORS 配置
    router.Use(cors.New(corsConfig))

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}
