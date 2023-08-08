package post

import (
	"nothing/internal/app/blog/model/common"
	"nothing/internal/app/blog/repository"
)

type FindReq struct {
	repository.BaseModel
	UserID             int64  `json:"user_id" form:"user_id" gorm:"column:user_id"`
	Type               []int  `json:"type" form:"type" gorm:"column:type"`
	Title              string `json:"title" form:"title" gorm:"column:title"`
	CategoryId         []int  `json:"category_id" form:"category_id" gorm:"column:category_id"`
	Tags               string `json:"tags" form:"tags" gorm:"column:tags"`
	Location           string `json:"location" form:"location" gorm:"column:location"`
	Order              int    `json:"order" form:"order"`
	Mode               string `json:"mode" form:"mode"`
	Page               common.Page
	RespAttachmentMode int `json:"resp_attachment_mode" form:"resp_attachment_mode"`
}
