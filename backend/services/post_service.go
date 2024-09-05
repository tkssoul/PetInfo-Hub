package services

import (
    "backend/models"
    "backend/repository"        
)

type PostCreation struct {
    Title   string `json:"title"`
    Summary string `json:"summary"`
    ThumbnailURL string `json:"thumbnail_url"`
    Tags string `json:"tags"`
    Views int `json:"views"`
    LikeCount int `json:"like_count"`    
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
}

type CommentCreation struct {
    UserID uint   `json:"user_id"`
    Text   string `json:"text"`
}

type PostService struct {
    postRepo    *repository.PostRepository
    commentRepo *repository.CommentRepository
}

func NewPostService(postRepo *repository.PostRepository, commentRepo *repository.CommentRepository) *PostService {
    return &PostService{
        postRepo:    postRepo,
        commentRepo: commentRepo,
    }
}

// CreatePost 创建动态
func (ps *PostService) CreatePost(postCreation PostCreation) error {
    newPost := models.Posts{
        Title: postCreation.Title,
        Summary: postCreation.Summary,
        ThumbnailURL: postCreation.ThumbnailURL,
        Tags: postCreation.Tags,
        Views: postCreation.Views,
        LikeCount: postCreation.LikeCount,
        User_ID: postCreation.UserID,
        Content: postCreation.Content,
    }

    return ps.postRepo.CreatePost(&newPost)
}

// GetPostByID 通过ID获取动态
func (ps *PostService) GetPostByID(postID uint) (*models.Posts, error) {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return nil, err
    }
    return post, nil
}

// UpdatePost 更新动态
func (ps *PostService) UpdatePost(postID uint, updatedPost models.Posts) error {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return err
    }

    // 更新数据
    post.Content = updatedPost.Content
    // 你可以在这里添加更多的字段更新
    // post.Title = updatedPost.Title
    // post.ImageURL = updatedPost.ImageURL

    return ps.postRepo.UpdatePost(post)
}

// DeletePost 删除动态
func (ps *PostService) DeletePost(postID uint) error {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return err
    }

    return ps.postRepo.DeletePost(post)
}

// LikePost 给动态点赞
func (ps *PostService) LikePost(postID uint) error {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return err
    }

    return ps.postRepo.IncrementLikes(post)
}

// 获取动态的点赞数
func (ps *PostService) GetLikesCount(postID uint) (int, error) {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return 0, err
    }

    return post.LikeCount, nil
}

// 取消点赞
func (ps *PostService) UnlikePost(postID uint) error {
    post, err := ps.postRepo.FindPostByID(postID)
    post.LikeCount--
    if err != nil {
        return err
    }
    return nil
}

// CommentOnPost 对动态评论
func (ps *PostService) CommentOnPost(postID uint, commentCreation CommentCreation) error {
    post, err := ps.postRepo.FindPostByID(postID)
    if err != nil {
        return err
    }

    newComment := models.Comments{
        Post_ID: post.ID,
        User_ID: commentCreation.UserID,
        Content: commentCreation.Text,
    }

    return ps.commentRepo.CreateComment(&newComment)
}

// 获取所有动态
func (ps *PostService) GetAllPosts() ([]models.Posts, error) {
    return ps.postRepo.GetAllPosts()
}