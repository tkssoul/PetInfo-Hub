package repository

import (    
    "gorm.io/gorm"
	"backend/models"
)

type CommentRepository struct {
    db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
    return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *models.Comment) error {
    result := r.db.Create(comment)
    return result.Error
}

func (r *CommentRepository) FindCommentsByPostID(postID uint) (*models.Comment, error) {
    var comments *models.Comment
    result := r.db.Where("post_id = ?", postID).Find(&comments)
    return comments, result.Error
}
func (r *CommentRepository) DeleteComment(comment *models.Comment) error {
    result := r.db.Delete(comment)
    return result.Error
}
