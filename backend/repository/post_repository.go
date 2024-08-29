package repository

import (
    "errors"
    "backend/models"
    "gorm.io/gorm"
)

type PostRepository struct {
    db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
    return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
    result := r.db.Create(post)
    return result.Error
}

func (r *PostRepository) FindPostByID(postID uint) (*models.Post, error) {
    var post models.Post
    result := r.db.First(&post, postID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该帖子")
        }
        return nil, result.Error
    }
    return &post, nil
}

func (r *PostRepository) IncrementLikes(post *models.Post) error {
    post.Likes++
    result := r.db.Save(post)
    return result.Error
}
