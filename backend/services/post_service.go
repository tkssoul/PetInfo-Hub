package services

import (    
    "backend/models"
    "backend/repository"
)

type PostCreation struct {
    UserID  uint   `json:"user_id"`
    Content string `json:"content"`
}

type CommentCreation struct {
    UserID uint   `json:"user_id"`
    Text   string `json:"text"`
}

func CreatePost(post PostCreation, repo *repository.PostRepository) error {
    newPost := models.Posts{
        User_ID:  post.UserID,
        Content: post.Content,
    }

    return repo.CreatePost(&newPost)
}

func LikePost(postID uint, repo *repository.PostRepository) error {
    post, err := repo.FindPostByID(postID)
    if err != nil {
        return err
    }

    return repo.IncrementLikes(post)
}

func CommentOnPost(postID uint, comment CommentCreation, commentRepo *repository.CommentRepository, postRepo *repository.PostRepository) error {
    post, err := postRepo.FindPostByID(postID)
    if err != nil {
        return err
    }

    newComment := models.Comments{
        Post_ID: post.Post_ID,
        User_ID: comment.UserID,
        Content:   comment.Text,
    }

    return commentRepo.CreateComment(&newComment)
}
