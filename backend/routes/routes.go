package routes

import (    
    "github.com/gin-gonic/gin"
    "backend/controllers" 
)

func SetupRoutes() *gin.Engine{
    r := gin.Default()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.POST("/posts", controllers.CreatePost)
    r.POST("/posts/{id:[0-9]+}/like", controllers.LikePost)
    r.POST("/posts/{id:[0-9]+}/comment", controllers.CommentOnPost)
    r.GET("/hotels", controllers.GetNearbyHotels)    
    
    return r
}
