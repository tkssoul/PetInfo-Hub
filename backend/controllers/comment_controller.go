package controllers

import (
    "net/http"    
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
)

type CommentController struct {
    commentService *services.CommentService
}

func NewCommentController(commentService *services.CommentService) *CommentController {
    return &CommentController{commentService: commentService}
}

// GetCommentsByPostID 获取特定动态的评论
func (cc *CommentController) GetCommentsByPostID(c *gin.Context) {    
	postIDStr := c.Param("post_id")	
	postIDUint, _ := strconv.ParseUint(postIDStr, 10, 64)
	postID := uint(postIDUint) 
    comments, err := cc.commentService.GetCommentsByPostID(postID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, comments)
}

// CreateComment 为特定动态添加评论
func (cc *CommentController) CreateComment(c *gin.Context) {
    var comment models.Comments
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := cc.commentService.CreateComment(&comment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "评论创建成功"})
}

// UpdateComment 更新特定评论
func (cc *CommentController) UpdateComment(c *gin.Context) {
    var comment models.Comments
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := cc.commentService.UpdateComment(&comment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "评论更新成功"})
}

// DeleteComment 删除特定评论
func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentIDStr := c.Param("comment_id")	
	commmentIDUint, _ := strconv.ParseUint(commentIDStr, 10, 64)
	commentID := uint(commmentIDUint) 
    err := cc.commentService.DeleteComment(commentID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "评论删除成功"})
}
