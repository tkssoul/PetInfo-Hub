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
// 创建评论
func (r *CommentRepository) CreateComment(comments *models.Comments) error {
    result := r.db.Create(comments)
    return result.Error
}

// 通过帖子ID查找评论
func (r *CommentRepository) FindCommentsByPostID(postID uint) (*models.Comments, error) {
    var comments *models.Comments
    result := r.db.Where("post_id = ?", postID).Find(&comments)
    return comments, result.Error
}

// 删除评论
func (r *CommentRepository) DeleteComment(comments *models.Comments) error {
    result := r.db.Delete(comments)
    return result.Error
}
