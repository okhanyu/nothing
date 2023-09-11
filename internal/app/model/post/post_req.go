package post

import (
	"nothing/internal/app/model/common"
	"nothing/internal/app/repository"
)

type FindReq struct {
	repository.BaseModel
	Fields             []string `json:"fields" form:"fields" gorm:"column:fields"`
	Filters            FindReqFilters
	Order              int         `json:"order" form:"order"`
	RespAttachmentMode int         `json:"resp_attachment_mode" form:"resp_attachment_mode"`
	Page               common.Page `json:"page" form:"page"`
	PartitionColumn    string      `json:"partition_column" form:"partition_column"`
}

type FindReqFilters struct {
	UserID      int64  `json:"user_id" form:"user_id" gorm:"column:user_id"`
	Type        []int  `json:"type" form:"type" gorm:"column:type"`
	Title       string `json:"title" form:"title" gorm:"column:title"`
	CategoryId  []int  `json:"category_id" form:"category_id" gorm:"column:category_id"`
	Tags        string `json:"tags" form:"tags" gorm:"column:tags"`
	Location    string `json:"location" form:"location" gorm:"column:location"`
	PrimaryType []int  `json:"primary_type" form:"primary_type"  gorm:"column:primary_type"`
	Hide        []int  `json:"hide" form:"hide"  gorm:"column:hide"`
	//Mode        string `json:"mode" form:"mode"`
}

type CreateReq struct {
	Main       PostMain         `json:"main" form:"main"`
	Attachment []PostAttachment `json:"attachment" form:"attachment"`
	Extend     PostExtend       `json:"extend" form:"extend"`
}
