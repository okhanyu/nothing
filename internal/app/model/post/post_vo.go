package post

import (
	"nothing/internal/app/model/user"
	"time"
)

type HiddenPostVo struct {
	ID        int64      `json:"id" form:"id"  `
	CreatedAt *time.Time `json:"created_at" form:"created_at"`
	Title     string     `json:"title" form:"title" `
}

type NormalPostVo struct {
	ID            int64                      `json:"id" form:"id"`
	CreatedAt     *time.Time                 `json:"created_at" form:"created_at"`
	UpdatedAt     *time.Time                 `json:"updated_at" form:"updated_at"`
	Type          int                        `json:"type" form:"type"`
	Title         string                     `json:"title" form:"title"`
	Summary       string                     `json:"summary" form:"summary"`
	Tags          string                     `json:"tags" form:"tags" `
	Location      string                     `json:"location" form:"location"`
	Category      PostCategoryVo             `json:"category" form:"category"`
	Extend        PostExtendVo               `json:"extend" form:"extend" `
	CategoryID    int                        `json:"category_id" form:"category_id" `
	UserID        int64                      `json:"user_id" form:"user_id"`
	User          user.UserVo                `json:"user" form:"user"`
	AttachmentMap map[int][]PostAttachmentVo `json:"attachment_map" form:"attachment_map"`
	Attachments   []PostAttachmentVo         `json:"attachments" form:"attachments"`
	Hide          int                        `json:"hide" form:"hide"`
}
type SimplePostVo struct {
	ID         int64          `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	CreatedAt  *time.Time     `json:"created_at" form:"created_at"  gorm:"column:created_at"`
	UpdatedAt  *time.Time     `json:"updated_at" form:"updated_at"  gorm:"column:updated_at"`
	Type       int            `json:"type" form:"type" gorm:"column:type"`
	Title      string         `json:"title" form:"title" gorm:"column:title"`
	Summary    string         `json:"summary" form:"summary" gorm:"column:summary"`
	Tags       string         `json:"tags" form:"tags" gorm:"column:tags"`
	Location   string         `json:"location" form:"location" gorm:"column:location"`
	CategoryID int            `json:"category_id" form:"category_id" `
	UserID     int64          `json:"user_id" form:"user_id"`
	Category   PostCategoryVo `json:"category" form:"category"`
	Extend     PostExtendVo   `json:"extend" form:"extend" `
	User       user.UserVo    `json:"user" form:"user"`
}

type PartitionPostVo struct {
	ID        int64      `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"  gorm:"column:created_at"`
	//UpdatedAt  *time.Time `json:"updated_at" form:"updated_at"  gorm:"column:updated_at"`
	Type         int    `json:"type" form:"type" gorm:"column:type"`
	Title        string `json:"title" form:"title" gorm:"column:title"`
	Summary      string `json:"summary" form:"summary" gorm:"column:summary"`
	Tags         string `json:"tags" form:"tags" gorm:"column:tags"`
	Location     string `json:"location" form:"location" gorm:"column:location"`
	CategoryID   int    `json:"category_id" form:"category_id" `
	CategoryName string `json:"category_name" form:"category_name" gorm:"column:category_name"`
	UserID       int64  `json:"user_id" form:"user_id"`
	//Category   PostCategoryVo `json:"category" form:"category"`
	//Extend     PostExtendVo   `json:"extend" form:"extend" `
	//User       user.UserVo    `json:"user" form:"user"`
}

type PostAttachmentVo struct {
	ID            int64      `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	CreatedAt     *time.Time `json:"created_at" form:"created_at"  gorm:"column:created_at"`
	PrimaryType   int        `json:"primary_type" form:"primary_type" gorm:"primary_type"`
	SecondaryType int        `json:"secondary_type" form:"secondary_type" gorm:"secondary_type"`
	Content       string     `json:"content" form:"content" gorm:"content"`
	PostID        int64      `json:"post_id" form:"post_id" gorm:"post_id"`
}

type PostExtendVo struct {
	ID    int64 `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	Likes int   `json:"likes" form:"likes" gorm:"column:likes"`
	Views int   `json:"views" form:"views" gorm:"column:views"`
}

type PostCategoryVo struct {
	ID   int64  `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	Name string `json:"name" form:"name" gorm:"name"`
}
