package post

import (
	"nothing/internal/app/repository"
)

type PostMain struct {
	repository.BaseModel
	UserID     int64  `json:"user_id" form:"user_id" gorm:"column:user_id"`
	Type       int    `json:"type" form:"type" gorm:"column:type"`
	Title      string `json:"title" form:"title" gorm:"column:title"`
	Summary    string `json:"summary" form:"summary" gorm:"column:summary"`
	CategoryID int    `json:"category_id" form:"category_id" gorm:"column:category_id"`
	Tags       string `json:"tags" form:"tags" gorm:"column:tags"`
	Location   string `json:"location" form:"location" gorm:"column:location"`
	Hide       int    `json:"hide" form:"hide" gorm:"column:hide"`
}

type PostAttachment struct {
	repository.BaseModel
	PrimaryType   int    `json:"primary_type" form:"primary_type" gorm:"primary_type"`
	SecondaryType int    `json:"secondary_type" form:"secondary_type" gorm:"secondary_type"`
	Content       string `json:"content" form:"content" gorm:"content"`
	PostID        int64  `json:"post_id" form:"post_id" gorm:"post_id"`
}

type PostExtend struct {
	ID    int64 `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	Likes int   `json:"likes" form:"likes" gorm:"column:likes"`
	Views int   `json:"views" form:"views" gorm:"column:views"`
}

type PostCategory struct {
	repository.BaseModel
	Name string `json:"name" form:"name" gorm:"name"`
}

func (PostMain) TableName() string {
	return repository.PostTable
}

func (PostAttachment) TableName() string {
	return repository.PostAttachmentTable
}

func (PostExtend) TableName() string {
	return repository.PostExtendTable
}

func (PostCategory) TableName() string {
	return repository.PostCategoryTable
}
