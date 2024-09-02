package repository

import (    
    "gorm.io/gorm"
	"backend/models"
    "errors"
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
func (r *CommentRepository) FindCommentsByPostID(postID uint) ([]models.Comments, error) {
    var comments []models.Comments
    result := r.db.Where("post_id = ?", postID).Find(&comments)
    return comments, result.Error
}

// 通过评论ID查找特定评论
func (r *CommentRepository) FindCommentByID(commentID uint) (*models.Comments, error) {
    var comment models.Comments
    result := r.db.First(&comment, commentID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("评论不存在")
        }
        return nil, result.Error
    }
    return &comment, nil
}

// 更新评论
func (r *CommentRepository) UpdateComment(comment *models.Comments) error {
    result := r.db.Save(comment)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("更新失败，评论不存在")
    }
    return nil
}

// 删除评论
func (r *CommentRepository) DeleteComment(comments *models.Comments) error {
    result := r.db.Delete(comments)
    return result.Error
}
