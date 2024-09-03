package controllers

import (
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "backend/models"
    "backend/services"
	"strconv"
)

type MessageController struct {
    messageService *services.MessageService
}

// NewMessageController 创建一个新的 MessageController 实例
func NewMessageController(messageService *services.MessageService) *MessageController {
    return &MessageController{messageService: messageService}
}

// CreateMessage 创建新的消息
func (mc *MessageController) CreateMessage(c *gin.Context) {
    var message models.Messages
    if err := c.ShouldBindJSON(&message); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    message.Send_at = time.Now() // 设置发送时间为当前时间

    if err := mc.messageService.CreateMessage(&message); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "消息创建成功"})
}

// GetMessageByID 获取特定消息的详细信息
func (mc *MessageController) GetMessageByID(c *gin.Context) {
    messageIDStr := c.Param("message_id")	
	messageIDUint, _ := strconv.ParseUint(messageIDStr, 10, 64)
	messageID := uint(messageIDUint) 
    msg, err := mc.messageService.GetMessageByID(messageID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, msg)
}

// GetMessagesByUserID 获取特定用户的消息列表
func (mc *MessageController) GetMessagesByUserID(c *gin.Context) {
    userIDStr := c.Param("userId")
	userIDUint,_ := strconv.ParseUint(userIDStr, 10, 64)
	userID := uint(userIDUint)
    messages, err := mc.messageService.GetMessagesByUserID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, messages)
}
