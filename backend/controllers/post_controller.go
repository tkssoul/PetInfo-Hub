package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
)

type PostController struct {
    postService *services.PostService
    commentService *services.CommentService
}

func NewPostController(postService *services.PostService, commentService *services.CommentService) *PostController {
    return &PostController{
        postService: postService,
        commentService: commentService,
    }
}

// CreatePost 创建动态
func (pc *PostController) CreatePost(c *gin.Context) {
    var postCreation services.PostCreation
    if err := c.ShouldBindJSON(&postCreation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.postService.CreatePost(postCreation)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "动态创建成功"})
}

// GetPostByID 通过ID获取动态
func (pc *PostController) GetPostByID(postID uint,c *gin.Context) {    
    post, err := pc.postService.GetPostByID(postID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

// UpdatePost 更新动态
func (pc *PostController) UpdatePost(postID uint,c *gin.Context) {

    var updatedPost models.Posts
    if err := c.ShouldBindJSON(&updatedPost); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.postService.UpdatePost(postID, updatedPost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "动态更新成功"})
}

// DeletePost 删除动态
func (pc *PostController) DeletePost(postID uint,c *gin.Context) {    
    err := pc.postService.DeletePost(postID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "动态删除成功"})
}

// LikePost 给动态点赞
func (pc *PostController) LikePost(postID uint,c *gin.Context) {    
    err := pc.postService.LikePost(postID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}

// CommentOnPost 对动态评论
func (pc *PostController) CommentOnPost(postID uint,c *gin.Context) {    
    var commentCreation services.CommentCreation
    if err := c.ShouldBindJSON(&commentCreation); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := pc.postService.CommentOnPost(postID, commentCreation)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "评论成功"})
}
