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

// 创建动态
func (r *PostRepository) CreatePost(post *models.Posts) error {
    result := r.db.Create(post)
    return result.Error
}

// 通过动态ID获取特定动态
func (r *PostRepository) FindPostByID(postID uint) (*models.Posts, error) {
    var post models.Posts
    result := r.db.First(&post, postID)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return nil, errors.New("找不到该帖子")
        }
        return nil, result.Error
    }
    return &post, nil
}

// 获取所有动态
func (r *PostRepository) FindAllPosts() ([]models.Posts, error) {
    var posts []models.Posts
    result := r.db.Find(&posts)
    return posts, result.Error
}

// 增加1个帖子的点赞数
func (r *PostRepository) IncrementLikes(posts *models.Posts) error {
    posts.Like_count++
    result := r.db.Save(posts)
    return result.Error
}

// 增加1个帖子的浏览数
func (r *PostRepository) IncrementViews(posts *models.Posts) error {
    posts.Views++
    result := r.db.Save(posts)
    return result.Error
}

// 减少1个帖子的点赞数
func (r *PostRepository) DecreaseLikes(posts *models.Posts) error {
    posts.Like_count--
    result := r.db.Save(posts)
    return result.Error
}

// 删除特定动态
func (r *PostRepository) DeletePost(post *models.Posts) error {
    result := r.db.Delete(post)
    return result.Error
}