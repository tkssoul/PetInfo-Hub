package controllers

import (
	"backend/models"
	"backend/repository"
	"backend/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var repo = repository.NewPostRepository(models.DB)

// 发表动态
func CreatePost(c *gin.Context) {
    var post services.PostCreation
    if err := c.BindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
    
    if err := services.CreatePost(post,repo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

// 点赞动态
func LikePost(c *gin.Context) {    
    postID_str := c.Param("id")
    postID_int, _ := strconv.Atoi(postID_str)
    postID := uint(postID_int)    

    if err := services.LikePost(postID,repo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

// 评论动态
func CommentOnPost(c *gin.Context) {    

    var request services.CommentCreateRequest
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    var postRepo = repository.NewCommentRepository(models.DB)
    if err := services.CreateComment(request,postRepo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}
