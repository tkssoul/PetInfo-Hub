package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
    "time"
)
var DB *gorm.DB
// 实名信息表
type ReanNameInfo struct {
    gorm.Model    
    Username string `gorm:"primaryKey" json:"real_name_id"`
    User_id  int    `gorm:"not null" json:"user_id"`
    Real_name string `gorm:"type:varchar(100) not null" json:"real_name"`

}

// 用户
type Users struct {
    gorm.Model
    User_id  int    `gorm:"primaryKey autoIncrement" json:"user_id"`
    Username string `gorm:"type:varchar(50) not null unique" json:"username"`
    Password string `gorm:"type:varchar(255) not null" json:"password"`
}

// 宠物
type Pets struct {
    gorm.Model
    Pet_id  int    `gorm:"primaryKey autoIncrement" json:"pet_id"`
    User_id int    `gorm:"not null" json:"user_id"`
    Name     string `gorm:"type:varchar(50) not null" json:"name"`
    Species string `gorm:"type:varchar(50) not null" json:"species"`
    Breed   string `gorm:"type:varchar(50) not null" json:"breed"`
    Age      int    `gorm:"type:int" json:"age"`
    Photo   string `gorm:"type:varchar(255)" json:"photo"`
}

// 动态
type Posts struct {
    gorm.Model    
    Post_id  int    `gorm:"primaryKey autoIncrement" json:"post_id"`
    User_id  int    `gorm:"not null" json:"user_id"`
    Title    string `gorm:"type:varchar(255) not null" json:"title"`
    Content  string `gorm:"type:text not null" json:"content"`
    Summary  string `gorm:"type:text not null" json:"summary"`
    Thumbnail_url string `gorm:"type:varchar(255)" json:"thumbnail_url"`
    Tags    string `gorm:"type:varchar(255)" json:"tags"`        
}

// 互动
type Interactions struct {
    Interaction_id int `gorm:"primaryKey autoIncrement" json:"interaction_id"`
    Post_id        int `gorm:"unique not null" json:"post_id"`
    User_id        int `gorm:"unique not null" json:"user_id"`
    Interaction_type string `gorm:"unique type:ENUM('like','comment') not null" json:"interaction_type"`
    Content string `gorm:"type:text" json:"content"`    
}

// 好友关系
type Friendship struct {
    gorm.Model
    Friendship_id int `gorm:"primaryKey autoIncrement" json:"friendship_id"`
    User_id       int `gorm:"unique not null" json:"user_id"`
    Friend_id     int `gorm:"unique not null" json:"friend_id"`    
}

// 消息
type Messages struct {
    gorm.Model
    Message_id int `gorm:"primaryKey autoIncrement" json:"message_id"`
    Sender_id   int `gorm:"unique not null" json:"sender_id"`
    Receiver_id int `gorm:"unique not null" json:"receiver_id"`
    Content    string `gorm:"type:text not null" json:"content"`
    Image_url string `gorm:"type:varchar(255)" json:"image_url"`
    Send_at     time.Time `json:"send_at"`
}

// 攻略
type Guide struct {
    Guide_id int `gorm:"primaryKey autoIncrement" json:"guide_id"`
    Title    string `gorm:"type:varchar(255) not null" json:"title"`
    Content  string `gorm:"type:text not null" json:"content"`
    Species  string `gorm:"type:varchar(100)" json:"species"`
    Age_range string `gorm:"type:varchar(50)" json:"age_range"`
    Category string `gorm:"type:varchar(100)" json:"category"`
    Author_id int `gorm:"unique not null" json:"author_id"`
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
    ShopID       int       `gorm:"primaryKey;autoIncrement" json:"shop_id"`  
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
    SitterID         int       `gorm:"primaryKey;autoIncrement" json:"sitter_id"`  
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

func InitDB() {
    var err error
    dsn := "root:rootpw1.@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }    
    DB.AutoMigrate(&ReanNameInfo{}, &Users{}, &Posts{}, &Pets{},&Interactions{},&Friendship{},&Messages{},&Guide{},&PetFriendlySpot{},&PetCareShop{},&PetSitter{},&PetBoardingDetail{})
}