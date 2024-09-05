package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
    "strconv"
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
    postCreation.Content = c.PostForm("content")
    postCreation.Summary = c.PostForm("summary")
    postCreation.Tags = c.PostForm("tags")
    postCreation.ThumbnailURL = c.PostForm("thumbnail_url")
    postCreation.Title = c.PostForm("title")
    userIDStr := c.Param("user_id")
    userIDUint, _ := strconv.ParseUint(userIDStr, 10, 64)
    postCreation.UserID = uint(userIDUint)    
    postCreation.Views = 0
    postCreation.LikeCount = 0
    err := pc.postService.CreatePost(postCreation)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "动态创建成功"})
}

// GetPostByID 通过ID获取动态
func (pc *PostController) GetPostByID(c *gin.Context) {    
    postIDStr := c.Param("post_id")	
	userIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
	postID := uint(userIDUint) 
    post, err := pc.postService.GetPostByID(postID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, post)
}

// UpdatePost 更新动态
func (pc *PostController) UpdatePost(c *gin.Context) {
    postIDStr := c.Param("pet_id")	
	postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
	postID := uint(postIDUint) 
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
func (pc *PostController) DeletePost(c *gin.Context) {    
    postIDStr := c.Param("pet_id")	
	postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
	postID := uint(postIDUint) 
    err := pc.postService.DeletePost(postID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "动态删除成功"})
}

// LikePost 给动态点赞
func (pc *PostController) LikePost(c *gin.Context) {    
    postIDStr := c.Param("pet_id")	
	postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
	postID := uint(postIDUint) 
    err := pc.postService.LikePost(postID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}

// 获取点赞数
func (pc *PostController) GetLikesCount(c *gin.Context) {
    postIDStr := c.Param("post_id")	
    postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
    postID := uint(postIDUint) 
    likesCount, err := pc.postService.GetLikesCount(postID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"likes_count": likesCount})
}

// 取消点赞
func (pc *PostController) DislikePost(c *gin.Context) {
    postIDStr := c.Param("post_id")	
    postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
    postID := uint(postIDUint) 
    err := pc.postService.UnlikePost(postID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "取消点赞成功"})
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

// 获取所有动态
func (pc *PostController) GetAllPosts(c *gin.Context) {
    posts, err := pc.postService.GetAllPosts()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, posts)
}