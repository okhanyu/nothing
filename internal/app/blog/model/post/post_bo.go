package post

import (
	"nothing/internal/app/blog/model/user"
	"nothing/internal/app/blog/repository"
)

const (
	Attachments = "Attachments"
	User        = "User"
	Extend      = "Extend"
	Category    = "Category"
)

type PostBo struct {
	repository.BaseModel
	UserID      int64            `json:"user_id" form:"user_id" gorm:"column:user_id"`
	Type        int              `json:"type" form:"type" gorm:"column:type"`
	Title       string           `json:"title" form:"title" gorm:"column:title"`
	Summary     string           `json:"summary" form:"summary" gorm:"column:summary"`
	CategoryID  int              `json:"category_id" form:"category_id" gorm:"column:category_id"`
	ExtendID    int64            `json:"extend_id" form:"extend_id" gorm:"column:extend_id"`
	Tags        string           `json:"tags" form:"tags" gorm:"column:tags"`
	Location    string           `json:"location" form:"location" gorm:"column:location"`
	Hide        int              `json:"hide" form:"hide" gorm:"column:hide"`
	Category    PostCategory     `json:"category" form:"category" gorm:"foreignKey:ID;references:CategoryID"`
	Extend      PostExtend       `json:"extend" form:"extend" gorm:"foreignKey:ID;references:ExtendID"`
	User        user.User        `json:"user" form:"user" gorm:"foreignKey:ID;references:UserID"`
	Attachments []PostAttachment `json:"attachments" form:"attachments" gorm:"foreignKey:PostID;references:ID"`
	// User        user.User        `json:"user" form:"user" gorm:"foreignKey:UserID"`
	// Attachments []PostAttachment `json:"attachments" form:"attachments" gorm:"foreignKey:PostID"`
}

type PostPartitionBo struct {
	repository.BaseModel
	UserID       int64        `json:"user_id" form:"user_id" gorm:"column:user_id"`
	Type         int          `json:"type" form:"type" gorm:"column:type"`
	Title        string       `json:"title" form:"title" gorm:"column:title"`
	Summary      string       `json:"summary" form:"summary" gorm:"column:summary"`
	CategoryID   int          `json:"category_id" form:"category_id" gorm:"column:category_id"`
	CategoryName string       `json:"category_name" form:"category_name" gorm:"column:category_name"`
	ExtendID     int64        `json:"extend_id" form:"extend_id" gorm:"column:extend_id"`
	Tags         string       `json:"tags" form:"tags" gorm:"column:tags"`
	Location     string       `json:"location" form:"location" gorm:"column:location"`
	Hide         int          `json:"hide" form:"hide" gorm:"column:hide"`
	Category     PostCategory `json:"category" form:"category" gorm:"foreignKey:ID;references:CategoryID"`
	Extend       PostExtend   `json:"extend" form:"extend" gorm:"foreignKey:ID;references:ExtendID"`
	User         user.User    `json:"user" form:"user" gorm:"foreignKey:ID;references:UserID"`
}

func (PostBo) TableName() string {
	return repository.PostTable
}
