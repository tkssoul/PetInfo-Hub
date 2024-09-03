package services

import (
    "errors"
    "backend/models"
    "backend/repository"
)

type CommentService struct {
    repo *repository.CommentRepository
}

func NewCommentService(repo *repository.CommentRepository) *CommentService {
    return &CommentService{repo: repo}
}

// 创建评论
func (s *CommentService) CreateComment(comment *models.Comments) error {
    // 这里可以添加更多业务逻辑，比如验证评论内容、用户权限等
    return s.repo.CreateComment(comment)
}

// 获取特定动态的评论
func (s *CommentService) GetCommentsByPostID(postID uint) ([]models.Comments, error) {
    comments, err := s.repo.FindCommentsByPostID(postID)
    if err != nil {
        return nil, err
    }
    if comments == nil {
        return nil, errors.New("未找到评论")
    }
    return comments, nil
}

// 更新特定评论
func (s *CommentService) UpdateComment(comment *models.Comments) error {
    // 确保评论存在
    existingComment, err := s.repo.FindCommentByID(comment.ID)
    if err != nil {
        return err
    }
    if existingComment == nil {
        return errors.New("评论不存在")
    }
    // 这里可以添加更多的业务逻辑，比如验证更新内容
    return s.repo.UpdateComment(comment)
}

// 删除特定评论
func (s *CommentService) DeleteComment(commentID uint) error {
    // 确保评论存在
    existingComment, err := s.repo.FindCommentByID(commentID)
    if err != nil {
        return err
    }
    if existingComment == nil {
        return errors.New("评论不存在")
    }
    return s.repo.DeleteComment(existingComment)
}
