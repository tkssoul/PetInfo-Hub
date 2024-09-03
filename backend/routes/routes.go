package routes

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
    "backend/repository"
    "backend/services"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
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

    // 动态、评论相关路由
    postRepo := repository.NewPostRepository(db)
    commentRepo := repository.NewCommentRepository(db)
    postService := services.NewPostService(postRepo,commentRepo)
    commentService := services.NewCommentService(commentRepo)
    postController := controllers.NewPostController(postService,commentService)
    commentController := controllers.NewCommentController(commentService)

    router.GET("/posts/:post_id/comments", commentController.GetCommentsByPostID)
    router.POST("/posts/:post_id/comments", commentController.CreateComment)
    router.PUT("/comments/:comment_id", commentController.UpdateComment)
    router.DELETE("/comments/:comment_id", commentController.DeleteComment)

    router.GET("/posts", postController.GetAllPosts)
    router.GET("/posts/:post_id", postController.GetPostByID)
    router.POST("/users/:user_id/posts", postController.CreatePost)
    router.PUT("/posts/:post_id", postController.UpdatePost)
    router.DELETE("/posts/:post_id", postController.DeletePost)

    // 点赞相关路由
    router.GET("/posts/:post_id/likes", postController.GetLikesCount)
    router.POST("/posts/:post_id/likes", postController.LikePost)
    router.PUT("/posts/:post_id/likes", postController.GetLikesCount)
    router.DELETE("/posts/:post_id/likes", postController.DislikePost)


    // 好友关系相关路由
    friendshipRepo := repository.NewFriendshipRepository(db)
    friendshipService := services.NewFriendshipService(friendshipRepo)
    friendshipController := controllers.NewFriendshipController(friendshipService)

    router.GET("/users/:user_id/friends", friendshipController.GetFriendsByUserID)
    router.POST("/users/:user_id/friends/:friend_id", friendshipController.AddFriend)
    router.DELETE("/users/:user_id/friends/:friend_id", friendshipController.RemoveFriend)

    // 消息相关路由
    messageRepo := repository.NewMessageRepository(db)
    messageService := services.NewMessageService(messageRepo)
    messageController := controllers.NewMessageController(messageService)

    router.GET("/users/:user_id/messages", messageController.GetMessagesByUserID)
    router.POST("/users/:user_id/messages", messageController.CreateMessage)

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
    spotService := services.NewPetFriendlySpotService(spotRepo)
    spotController := controllers.NewPetFriendlySpotController(spotService)

    router.GET("/pet-friendly-spots", spotController.GetAllPetFriendlySpots)
    router.GET("/pet-friendly-spots/:spot_id", spotController.GetPetFriendlySpotByID)
    router.POST("/create-pets-friendly-spot", spotController.CreatePetFriendlySpot)
    router.PUT("/update-pets-friendly-spot/:spot_id", spotController.UpdatePetFriendlySpot)
    router.DELETE("/delete-pets-friendly-spot/:spot_id", spotController.DeletePetFriendlySpot)

    // 服务店铺相关路由
    petCareShopRepo := repository.NewPetCareShopRepository(db)
    petCareShopService := services.NewPetCareShopService(petCareShopRepo)
    petCareShopController := controllers.NewPetCareShopController(petCareShopService)

    router.GET("/pet-care-shops", petCareShopController.GetAllPetCareShops)
    router.GET("/pet-care-shops/:shop_id", petCareShopController.GetPetCareShopByID)
    router.POST("/create-pet-care-shops", petCareShopController.CreatePetCareShop)
    router.PUT("/update-pet-care-shops/:shop_id", petCareShopController.UpdatePetCareShop)
    router.DELETE("/delete-pet-care-shops/:shop_id", petCareShopController.DeletePetCareShop)

    // 寄养人相关路由
    petSitterRepo := repository.NewPetSitterRepository(db)
    petSitterService := services.NewPetSitterService(petSitterRepo)
    petSitterController := controllers.NewPetSitterController(petSitterService)

    router.GET("/pet-sitters", petSitterController.GetAllPetSitters)
    router.GET("/pet-sitters/:sitter_id", petSitterController.GetPetSitterByID)

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
