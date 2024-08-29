package services

import ( 
    "backend/models"
    "backend/repository"
)

type CommentCreateRequest struct {
    PostID  uint   `json:"post_id"`
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
}

type CommentUpdateRequest struct {
    Content string `json:"content"`
}

type CommentResponse struct {
    ID      uint   `json:"id"`
    PostID  uint   `json:"post_id"`
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
}

// 发表评论
func CreateComment(request CommentCreateRequest, commentRepo *repository.CommentRepository) (error) {
    // Create a new comment instance
    newComment := models.Comment{
        PostID:  request.PostID,
        UserID:  request.UserID,
        Text: request.Content,
    }

    // Save the comment in the database
    err := commentRepo.CreateComment(&newComment)
    if err != nil {
        return  err
    }

    return  nil
}

// 定位评论
func GetCommentsByPostID(postID uint, commentRepo *repository.CommentRepository) (*models.Comment, error) {
    comments, err := commentRepo.FindCommentsByPostID(postID)
    if err != nil {
        return nil, err
    }

    return comments, nil
}

// 删除评论
func DeleteComment(commentID uint, commentRepo *repository.CommentRepository) error {
    // Find the comment
    comment, err := commentRepo.FindCommentsByPostID(commentID)
    if err != nil {
        return err
    }

    // Delete the comment
    err = commentRepo.DeleteComment(comment)
    if err != nil {
        return err
    }

    return nil
}
