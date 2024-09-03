package models

import (
	"time"	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

// 实名信息
type RealName struct {
    gorm.Model
    User_ID  uint    `gorm:"primaryKey autoIncrement" json:"user_id"`
    RIDNum string `gorm:"unique" json:"rid_num"`         
    Real_name string `gorm:"type:varchar(100) not null" json:"real_name"`
}

// 用户
type Users struct {
    gorm.Model
    User_ID  uint    `gorm:"primaryKey autoIncrement" json:"user_id"`
    Username string `gorm:"type:varchar(50) not null unique" json:"username"`
    Password string `gorm:"type:varchar(255) not null" json:"password"`        
}

// 宠物
type Pets struct {
    gorm.Model
    Pet_ID  uint    `gorm:"primaryKey autoIncrement" json:"pet_id"`
    User_ID int    `gorm:"not null" json:"user_id"`
    Name     string `gorm:"type:varchar(50) not null" json:"name"`
    Species string `gorm:"type:varchar(50) not null" json:"species"`
    Breed   string `gorm:"type:varchar(50) not null" json:"breed"`
    Age      int    `gorm:"type:int" json:"age"`
    Photo   string `gorm:"type:varchar(255)" json:"photo"`
}

// 动态
type Posts struct {
    gorm.Model    
    Post_ID  uint    `gorm:"primaryKey autoIncrement" json:"post_id"`
    User_ID  uint    `gorm:"not null" json:"user_id"`
    Title    string `gorm:"type:varchar(255) not null" json:"title"`
    Content  string `gorm:"type:text not null" json:"content"`
    Summary  string `gorm:"type:text not null" json:"summary"`
    Thumbnail_url string `gorm:"type:varchar(255)" json:"thumbnail_url"`
    Tags    string `gorm:"type:varchar(255)" json:"tags"`      
    Views  int    `gorm:"type:int" json:"views"`
    Like_count int `gorm:"type:int" json:"like_count"`  
}

// 点赞
type Likes struct {
    gorm.Model
    Like_ID int `gorm:"primaryKey autoIncrement" json:"like_id"`
    User_ID int `gorm:"unique not null" json:"user_id"`
    Post_ID int `gorm:"unique not null" json:"post_id"`
}

// 评论 
type Comments struct {
    gorm.Model
    Comment_ID uint `gorm:"primaryKey autoIncrement" json:"comment_id"`
    User_ID    uint `gorm:"unique not null" json:"user_id"`
    Post_ID    uint `gorm:"unique not null" json:"post_id"`
    Parent_Comment_ID int `gorm:"unique not null" json:"parent_comment_id"`
    Content    string `gorm:"type:text not null" json:"content"`    
}

// 好友关系
type Friendship struct {
    gorm.Model
    Friendship_ID int `gorm:"primaryKey autoIncrement" json:"friendship_id"`
    User_ID       int `gorm:"unique not null" json:"user_id"`
    Friend_ID     int `gorm:"unique not null" json:"friend_id"`    
}

// 消息
type Messages struct {
    gorm.Model
    Message_ID uint `gorm:"primaryKey autoIncrement" json:"message_id"`
    Sender_ID   uint `gorm:"unique not null" json:"sender_id"`
    Receiver_ID uint `gorm:"unique not null" json:"receiver_id"`
    Content    string `gorm:"type:text not null" json:"content"`
    Image_url string `gorm:"type:varchar(255)" json:"image_url"`
    Send_at     time.Time `json:"send_at"`
}

// 攻略
type Guide struct {
    Guide_ID uint `gorm:"primaryKey autoIncrement" json:"guide_id"`
    Title    string `gorm:"type:varchar(255) not null" json:"title"`
    Content  string `gorm:"type:text not null" json:"content"`
    Species  string `gorm:"type:varchar(100)" json:"species"`
    Age_range string `gorm:"type:varchar(50)" json:"age_range"`
    Category string `gorm:"type:varchar(100)" json:"category"`
    Author_ID int `gorm:"unique not null" json:"author_id"`
    Image_url string `gorm:"type:varchar(255)" json:"image_url"`
}

// 景点信息
type PetFriendlySpot struct {  
    gorm.Model  
    SpotID         uint      `gorm:"primaryKey;autoIncrement" json:"spot_id"`  
    Name           string    `gorm:"size:255;not null" json:"name"`  
    Location       string    `gorm:"size:255;not null" json:"location"`  
    Description    string    `gorm:"type:text" json:"description"`  
    OpeningHours   string    `gorm:"size:100" json:"opening_hours"`  
    EntryFee       string    `gorm:"size:50" json:"entry_fee"`  
    ContactInfo    string    `gorm:"size:255" json:"contact_info"`  
    Rating         float64   `gorm:"default:0" json:"rating"`  
    Tags           string    `gorm:"size:255" json:"tags"`  
    PetActivities  string    `gorm:"type:text" json:"pet_activities"`  
}  

// 服务信息
type PetCareShop struct {  
    gorm.Model  
    Shop_ID       uint       `gorm:"primaryKey;autoIncrement" json:"shop_id"`  
    Name         string    `gorm:"size:255;not null" json:"name"`  
    Location     string    `gorm:"size:255;not null" json:"location"`  
    Description  string    `gorm:"type:text" json:"description"`  
    Services     string    `gorm:"size:255" json:"services"`  
    ContactInfo  string    `gorm:"size:255" json:"contact_info"`  
    OpeningHours string    `gorm:"size:100" json:"opening_hours"`  
    Rating       float64   `gorm:"default:0" json:"rating"`  
    Tags         string    `gorm:"size:255" json:"tags"`     
}  

// 寄养人
type PetSitter struct {  
    gorm.Model  
    SitterID         uint       `gorm:"primaryKey;autoIncrement" json:"sitter_id"`  
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
    BoardingID     int       `gorm:"primaryKey;autoIncrement" json:"boarding_id"`  
    SitterID       int       `gorm:"not null" json:"sitter_id"`  
    PetName        string    `gorm:"size:255" json:"pet_name"`  
    PetType        string    `gorm:"size:50" json:"pet_type"`  
    StartDate      time.Time `gorm:"type:date" json:"start_date"`  
    EndDate        time.Time `gorm:"type:date" json:"end_date"`  
    SpecialRequirements string `gorm:"type:text" json:"special_requirements"`  
}  

func InitDB() *gorm.DB{
    var err error
    dsn := "root:666.@tcp(127.0.0.1:3306)/PetDB?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("连接数据库失败, error=" + err.Error())
    }    
    DB.AutoMigrate(&Users{}, &Posts{}, &Pets{},&Likes{},&Comments{},&Friendship{},&Messages{},&Guide{},&PetFriendlySpot{},&PetCareShop{},&PetSitter{},&PetBoardingDetail{})
    DB.Create(&Users{Username: "admin", Password: "adminpw",User_ID:1})
    return DB
}