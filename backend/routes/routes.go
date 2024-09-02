package main

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
    "backend/repository"
    "backend/services"
    "gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // 用户相关路由
    userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    router.GET("/users", userController.GetAllUsers)
    router.GET("/users/:user_id", userController.GetUserByID)
    router.POST("/users", userController.CreateUser)
    router.PUT("/users/:user_id", userController.UpdateUser)
    router.DELETE("/users/:user_id", userController.DeleteUser)

    router.GET("/users/:user_id/real-name-info", userController.GetRealNameInfo)
    router.POST("/users/:user_id/real-name-info", userController.CreateRealNameInfo)
    router.PUT("/users/:user_id/real-name-info", userController.UpdateRealNameInfo)
    router.DELETE("/users/:user_id/real-name-info", userController.DeleteRealNameInfo)

    // 宠物相关路由
    petRepo := repository.NewPetRepository(db)
    petService := services.NewPetService(petRepo)
    petController := controllers.NewPetController(petService)

    router.GET("/users/:user_id/pets", petController.GetPetsByUserID)
    router.GET("/pets/:pet_id", petController.GetPetByID)
    router.POST("/users/:user_id/pets", petController.CreatePet)
    router.PUT("/pets/:pet_id", petController.UpdatePet)
    router.DELETE("/pets/:pet_id", petController.DeletePet)

    // 动态相关路由
    postRepo := repository.NewPostRepository(db)
    postService := services.NewPostService(postRepo)
    postController := controllers.NewPostController(postService)

    router.GET("/posts", postController.GetAllPosts)
    router.GET("/posts/:post_id", postController.GetPostByID)
    router.POST("/users/:user_id/posts", postController.CreatePost)
    router.PUT("/posts/:post_id", postController.UpdatePost)
    router.DELETE("/posts/:post_id", postController.DeletePost)

    // 点赞相关路由
    router.GET("/posts/:post_id/likes", postController.GetLikesCount)
    router.POST("/posts/:post_id/likes", postController.LikePost)
    router.PUT("/posts/:post_id/likes", postController.UpdateLikesCount)
    router.DELETE("/posts/:post_id/likes", postController.DislikePost)

    // 评论相关路由
    commentRepo := repository.NewCommentRepository(db)
    commentService := services.NewCommentService(commentRepo)
    commentController := controllers.NewCommentController(commentService)

    router.GET("/posts/:post_id/comments", commentController.GetCommentsByPostID)
    router.POST("/posts/:post_id/comments", commentController.CreateComment)
    router.PUT("/comments/:comment_id", commentController.UpdateComment)
    router.DELETE("/comments/:comment_id", commentController.DeleteComment)

    // 好友关系相关路由
    friendshipRepo := repository.NewFriendshipRepository(db)
    friendshipService := services.NewFriendshipService(friendshipRepo, userRepo)
    friendshipController := controllers.NewFriendshipController(friendshipService)

    router.GET("/users/:user_id/friends", friendshipController.GetFriendsByUserID)
    router.POST("/users/:user_id/friends/:friend_id", friendshipController.AddFriend)
    router.DELETE("/users/:user_id/friends/:friend_id", friendshipController.RemoveFriend)

    // 消息相关路由
    messageRepo := repository.NewMessageRepository(db)
    messageService := services.NewMessageService(messageRepo)
    messageController := controllers.NewMessageController(messageService)

    router.GET("/users/:user_id/messages", messageController.GetMessagesByUserID)
    router.POST("/users/:receiver_id/messages", messageController.SendMessage)

    // 攻略相关路由
    guideRepo := repository.NewGuideRepository(db)
    guideService := services.NewGuideService(guideRepo)
    guideController := controllers.NewGuideController(guideService)

    router.GET("/guides", guideController.GetAllGuides)
    router.GET("/guides/:guide_id", guideController.GetGuideByID)
    router.POST("/users/:user_id/guides", guideController.CreateGuide)
    router.PUT("/guides/:guide_id", guideController.UpdateGuide)
    router.DELETE("/guides/:guide_id", guideController.DeleteGuide)

    // 景点相关路由
    spotRepo := repository.NewSpotRepository(db)
    spotService := services.NewSpotService(spotRepo)
    spotController := controllers.NewSpotController(spotService)

    router.GET("/pet-friendly-spots", spotController.GetAllSpots)
    router.GET("/pet-friendly-spots/:spot_id", spotController.GetSpotByID)
    router.POST("/create-pets-friendly-spot", spotController.CreateSpot)
    router.PUT("/update-pets-friendly-spot/:spot_id", spotController.UpdateSpot)
    router.DELETE("/delete-pets-friendly-spot/:spot_id", spotController.DeleteSpot)

    // 服务店铺相关路由
    petCareShopRepo := repository.NewPetCareShopRepository(db)
    petCareShopService := services.NewPetCareShopService(petCareShopRepo)
    petCareShopController := controllers.NewPetCareShopController(petCareShopService)

    router.GET("/pet-care-shops", petCareShopController.GetAllShops)
    router.GET("/pet-care-shops/:shop_id", petCareShopController.GetShopByID)
    router.POST("/create-pet-care-shops", petCareShopController.CreateShop)
    router.PUT("/update-pet-care-shops/:shop_id", petCareShopController.UpdateShop)
    router.DELETE("/delete-pet-care-shops/:shop_id", petCareShopController.DeleteShop)

    // 寄养人相关路由
    petSitterRepo := repository.NewPetSitterRepository(db)
    petSitterService := services.NewPetSitterService(petSitterRepo)
    petSitterController := controllers.NewPetSitterController(petSitterService)

    router.GET("/pet-sitters", petSitterController.GetAllSitters)
    router.GET("/pet-sitters/:sitter_id", petSitterController.GetSitterByID)

    // 寄养信息相关路由
    petBoardingDetailRepo := repository.NewPetBoardingDetailRepository(db)
    petBoardingDetailService := services.NewPetBoardingDetailService(petBoardingDetailRepo)
    petBoardingDetailController := controllers.NewPetBoardingDetailController(petBoardingDetailService)

    router.GET("/pet-sitters/:sitter_id/boarding-details", petBoardingDetailController.GetBoardingDetailsBySitterID)
    router.GET("/pet-boarding-details/:boarding_id", petBoardingDetailController.GetBoardingDetailByID)
    router.POST("/pet-sitters/:sitter_id/boarding-details", petBoardingDetailController.CreateBoardingDetail)
    router.PUT("/pet-boarding-details/:boarding_id", petBoardingDetailController.UpdateBoardingDetail)
    router.DELETE("/pet-boarding-details/:boarding_id", petBoardingDetailController.DeleteBoardingDetail)

    return router
}
