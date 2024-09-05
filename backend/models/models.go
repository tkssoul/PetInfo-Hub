package models

import (
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户
type Users struct {
    gorm.Model    
    Username     string      `gorm:"type:varchar(50) not null unique" json:"username" form:"username"`
    Password     string      `gorm:"type:varchar(255) not null" json:"password" form:"password"`
    RealName     RealName    `gorm:"foreignKey:User_ID" json:"real_name"`
    Pets        []Pets      `gorm:"foreignKey:User_ID" json:"pets"`
    Posts       []Posts     `gorm:"foreignKey:User_ID" json:"posts"`
    Messages    []Messages  `gorm:"foreignKey:Sender_ID" json:"messages"`
}

// 实名信息
type RealName struct {
    gorm.Model        
    RIDNum     string `gorm:"unique" json:"rid_num"`         
    RealName   string `gorm:"type:varchar(100) not null" json:"real_name"`    
    User_ID    uint   `gorm:"not null" json:"user_id"` // 外键
}

// 宠物
type Pets struct {
    gorm.Model    
    User_ID    uint   `gorm:"not null" json:"user_id"` // 外键
    Name      string `gorm:"type:varchar(50) not null" json:"name"`
    Species   string `gorm:"type:varchar(50) not null" json:"species"`
    Breed     string `gorm:"type:varchar(50) not null" json:"breed"`
    Age       int    `gorm:"type:int" json:"age"`
    Photo     string `gorm:"type:varchar(255)" json:"photo"`
}

// 动态
type Posts struct {
    gorm.Model        
    User_ID         uint      `gorm:"not null" json:"user_id"`  // 外键
    Title          string    `gorm:"type:varchar(255) not null" json:"title"`
    Content        string    `gorm:"type:text not null" json:"content"`
    Summary        string    `gorm:"type:text not null" json:"summary"`
    ThumbnailURL   string    `gorm:"type:varchar(255)" json:"thumbnail_url"`
    Tags           string    `gorm:"type:varchar(255)" json:"tags"`      
    Views          int       `gorm:"type:int" json:"views"`
    LikeCount      int       `gorm:"type:int" json:"like_count"`      
    Likes          []Likes   `gorm:"foreignKey:Post_ID" json:"likes"`
    Comments       []Comments `gorm:"foreignKey:Post_ID" json:"comments"`
}

// 点赞
type Likes struct {
    gorm.Model    
    User_ID  uint `gorm:"not null" json:"user_id"`  // 外键
    Post_ID  uint `gorm:"not null" json:"post_id"`  // 外键
}

// 评论
type Comments struct {
    gorm.Model    
    User_ID         uint `gorm:"not null" json:"user_id"`  // 外键
    Post_ID         uint `gorm:"not null" json:"post_id"`  // 外键
    ParentComment_ID uint `gorm:"" json:"parent_comment_id"` // 父评论ID，可为空
    Content        string `gorm:"type:text not null" json:"content"`    
}

// 好友关系
type Friendship struct {
    gorm.Model    
    User_ID        uint `gorm:"not null" json:"user_id"`      // 外键
    Friend_ID      uint `gorm:"not null" json:"friend_id"`    // 外键
}

// 消息
type Messages struct {
    gorm.Model    
    Sender_ID     uint `gorm:"not null" json:"sender_id"`   // 外键
    Receiver_ID   uint `gorm:"not null" json:"receiver_id"` // 外键
    Content      string `gorm:"type:text not null" json:"content"`
    ImageURL     string `gorm:"type:varchar(255)" json:"image_url"`
    SendAt       time.Time `json:"send_at"`
}

// 攻略
type Guide struct {
    gorm.Model    
    Title        string `gorm:"type:varchar(255) not null" json:"title"`
    Content      string `gorm:"type:text not null" json:"content"`
    Species      string `gorm:"type:varchar(100)" json:"species"`
    AgeRange     string `gorm:"type:varchar(50)" json:"age_range"`
    Category     string `gorm:"type:varchar(100)" json:"category"`
    AuthorID     uint   `gorm:"not null" json:"author_id"`  // 外键
    ImageURL     string `gorm:"type:varchar(255)" json:"image_url"`
}

// 景点信息
type PetFriendlySpot struct {  
    gorm.Model      
    Name           string    `gorm:"size:255;not null" json:"name"`  
    Location       string    `gorm:"size:255;not null" json:"location"`  
    Description    string    `gorm:"type:text" json:"description"`  
    OpeningHours   string    `gorm:"size:100" json:"opening_hours"`  
    EntryFee       string    `gorm:"size:50" json:"entry_fee"`  
    ContactInfo    string    `gorm:"size:255" json:"contact_info"`  
    Rating         float64   `gorm:"default:0" json:"rating"`  
    Tags           string    `gorm:"size:255" json:"tags"`  
    PetActivities  string    `gorm:"type:text" json:"pet_activities"`  
    ImgURL         string    `gorm:"size:255" json:"img_url"`
}  

// 服务信息
type PetCareShop struct {  
    gorm.Model      
    Name         string    `gorm:"size:255;not null" json:"name"`  
    Location     string    `gorm:"size:255;not null" json:"location"`  
    Description  string    `gorm:"type:text" json:"description"`  
    Services     string    `gorm:"size:255" json:"services"`  
    ContactInfo  string    `gorm:"size:255" json:"contact_info"`  
    OpeningHours string    `gorm:"size:100" json:"opening_hours"`  
    Rating       float64   `gorm:"default:0" json:"rating"`  
    Tags         string    `gorm:"size:255" json:"tags"` 
    ImgURL       string    `gorm:"size:255" json:"img_url"`    
}  

// 寄养人
type PetSitter struct {  
    gorm.Model       
    Name             string    `gorm:"size:255;not null" json:"name"`  
    ExperienceYears  int       `gorm:"not null" json:"experience_years"`  
    PetTypes         string    `gorm:"size:255" json:"pet_types"`  
    ProfilePicture   string    `gorm:"size:255" json:"profile_picture"`  
    EnvironmentPicture string  `gorm:"size:255" json:"environment_picture"`  
    Bio              string    `gorm:"type:text" json:"bio"`  
    ContactInfo      string    `gorm:"size:255" json:"contact_info"`      
}  

// 寄养信息
type PetBoardingDetail struct {  
    gorm.Model        
    Sitter_ID       uint       `gorm:"not null" json:"sitter_id"`  // 外键
    PetName        string    `gorm:"size:255" json:"pet_name"`  
    PetType        string    `gorm:"size:50" json:"pet_type"`  
    StartDate      time.Time `gorm:"type:date" json:"start_date"`  
    EndDate        time.Time `gorm:"type:date" json:"end_date"`  
    SpecialRequirements string `gorm:"type:text" json:"special_requirements"`  
}  



func InitDB() *gorm.DB{    
    dsn := "root:666@tcp(127.0.0.1:3306)/PetDB?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("连接数据库失败, error=" + err.Error())
    }    
    DB.AutoMigrate(&Users{},&RealName{}, &Posts{}, &Pets{},&Likes{},&Comments{},&Friendship{},&Messages{},&Guide{},&PetFriendlySpot{},&PetCareShop{},&PetSitter{},&PetBoardingDetail{})
    return DB
}