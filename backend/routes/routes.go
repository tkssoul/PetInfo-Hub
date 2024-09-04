package routes

import (
    "github.com/gin-gonic/gin"
    "backend/controllers"
    "backend/repository"
    "backend/services"
    "gorm.io/gorm"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "fmt"
    "net/http"
)

// JWT
var (
	secretKey = []byte("+onEiSrXvOwVChnWmt2JYGaFoP15eUc/wt4H+98zYmM=") // 保密并安全存储
)

// Claims 定义JWT的Claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(c *gin.Context) {
    username := c.PostForm("username")
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 设置2小时token过期
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token是必需项"})
			c.Abort()
			return
		}
        // 如果包含Bearer前缀，需要去除
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "非法的token"})
			c.Abort()
			return
		}

		c.Next()
	}
}


func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    router.GET("/login", GenerateToken)
    

    protected := router.Group("/")
    protected.Use(AuthMiddleware())

    // 用户相关路由
    userRepo := repository.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)
    // 用于注册
    router.POST("/users", userController.CreateUser)
    {
    protected.GET("/users", userController.GetAllUsers)
    protected.GET("/users/:user_id", userController.GetUserByID)    
    protected.PUT("/users/:user_id", userController.UpdateUser)
    protected.DELETE("/users/:user_id", userController.DeleteUser)

    protected.GET("/users/:user_id/real-name-info", userController.GetRealNameInfo)
    protected.POST("/users/:user_id/real-name-info", userController.CreateRealNameInfo)
    protected.PUT("/users/:user_id/real-name-info", userController.UpdateRealNameInfo)
    protected.DELETE("/users/:user_id/real-name-info", userController.DeleteRealNameInfo)

    // 宠物相关路由
    petRepo := repository.NewPetRepository(db)
    petService := services.NewPetService(petRepo)
    petController := controllers.NewPetController(petService)

    protected.GET("/users/:user_id/pets", petController.GetPetsByUserID)
    protected.GET("/pets/:pet_id", petController.GetPetByID)
    protected.POST("/users/:user_id/pets", petController.CreatePet)
    protected.PUT("/pets/:pet_id", petController.UpdatePet)
    protected.DELETE("/pets/:pet_id", petController.DeletePet)

    // 动态、评论相关路由
    postRepo := repository.NewPostRepository(db)
    commentRepo := repository.NewCommentRepository(db)
    postService := services.NewPostService(postRepo,commentRepo)
    commentService := services.NewCommentService(commentRepo)
    postController := controllers.NewPostController(postService,commentService)
    commentController := controllers.NewCommentController(commentService)

    protected.GET("/posts/:post_id/comments", commentController.GetCommentsByPostID)
    protected.POST("/posts/:post_id/comments", commentController.CreateComment)
    protected.PUT("/comments/:comment_id", commentController.UpdateComment)
    protected.DELETE("/comments/:comment_id", commentController.DeleteComment)

    protected.GET("/posts", postController.GetAllPosts)
    protected.GET("/posts/:post_id", postController.GetPostByID)
    protected.POST("/users/:user_id/posts", postController.CreatePost)
    protected.PUT("/posts/:post_id", postController.UpdatePost)
    protected.DELETE("/posts/:post_id", postController.DeletePost)

    // 点赞相关路由
    protected.GET("/posts/:post_id/likes", postController.GetLikesCount)
    protected.POST("/posts/:post_id/likes", postController.LikePost)
    protected.PUT("/posts/:post_id/likes", postController.GetLikesCount)
    protected.DELETE("/posts/:post_id/likes", postController.DislikePost)


    // 好友关系相关路由
    friendshipRepo := repository.NewFriendshipRepository(db)
    friendshipService := services.NewFriendshipService(friendshipRepo)
    friendshipController := controllers.NewFriendshipController(friendshipService)

    protected.GET("/users/:user_id/friends", friendshipController.GetFriendsByUserID)
    protected.POST("/users/:user_id/friends/:friend_id", friendshipController.AddFriend)
    protected.DELETE("/users/:user_id/friends/:friend_id", friendshipController.RemoveFriend)

    // 消息相关路由
    messageRepo := repository.NewMessageRepository(db)
    messageService := services.NewMessageService(messageRepo)
    messageController := controllers.NewMessageController(messageService)

    protected.GET("/users/:user_id/messages", messageController.GetMessagesByUserID)
    protected.POST("/users/:user_id/messages", messageController.CreateMessage)

    // 攻略相关路由
    guideRepo := repository.NewGuideRepository(db)
    guideService := services.NewGuideService(guideRepo)
    guideController := controllers.NewGuideController(guideService)

    protected.GET("/guides", guideController.GetAllGuides)
    protected.GET("/guides/:guide_id", guideController.GetGuideByID)
    protected.POST("/users/:user_id/guides", guideController.CreateGuide)
    protected.PUT("/guides/:guide_id", guideController.UpdateGuide)
    protected.DELETE("/guides/:guide_id", guideController.DeleteGuide)

    // 景点相关路由
    spotRepo := repository.NewSpotRepository(db)
    spotService := services.NewPetFriendlySpotService(spotRepo)
    spotController := controllers.NewPetFriendlySpotController(spotService)

    protected.GET("/pet-friendly-spots", spotController.GetAllPetFriendlySpots)
    protected.GET("/pet-friendly-spots/:spot_id", spotController.GetPetFriendlySpotByID)
    protected.POST("/create-pets-friendly-spot", spotController.CreatePetFriendlySpot)
    protected.PUT("/update-pets-friendly-spot/:spot_id", spotController.UpdatePetFriendlySpot)
    protected.DELETE("/delete-pets-friendly-spot/:spot_id", spotController.DeletePetFriendlySpot)

    // 服务店铺相关路由
    petCareShopRepo := repository.NewPetCareShopRepository(db)
    petCareShopService := services.NewPetCareShopService(petCareShopRepo)
    petCareShopController := controllers.NewPetCareShopController(petCareShopService)

    protected.GET("/pet-care-shops", petCareShopController.GetAllPetCareShops)
    protected.GET("/pet-care-shops/:shop_id", petCareShopController.GetPetCareShopByID)
    protected.POST("/create-pet-care-shops", petCareShopController.CreatePetCareShop)
    protected.PUT("/update-pet-care-shops/:shop_id", petCareShopController.UpdatePetCareShop)
    protected.DELETE("/delete-pet-care-shops/:shop_id", petCareShopController.DeletePetCareShop)

    // 寄养人相关路由
    petSitterRepo := repository.NewPetSitterRepository(db)
    petSitterService := services.NewPetSitterService(petSitterRepo)
    petSitterController := controllers.NewPetSitterController(petSitterService)

    protected.GET("/pet-sitters", petSitterController.GetAllPetSitters)
    protected.GET("/pet-sitters/:sitter_id", petSitterController.GetPetSitterByID)

    // 寄养信息相关路由
    petBoardingDetailRepo := repository.NewPetBoardingDetailRepository(db)
    petBoardingDetailService := services.NewPetBoardingDetailService(petBoardingDetailRepo)
    petBoardingDetailController := controllers.NewPetBoardingDetailController(petBoardingDetailService)

    protected.GET("/pet-sitters/:sitter_id/boarding-details", petBoardingDetailController.GetBoardingDetailsBySitterID)
    protected.GET("/pet-boarding-details/:boarding_id", petBoardingDetailController.GetBoardingDetailByID)
    protected.POST("/pet-sitters/:sitter_id/boarding-details", petBoardingDetailController.CreateBoardingDetail)
    protected.PUT("/pet-boarding-details/:boarding_id", petBoardingDetailController.UpdateBoardingDetail)
    protected.DELETE("/pet-boarding-details/:boarding_id", petBoardingDetailController.DeleteBoardingDetail)
    }

    return router
}
