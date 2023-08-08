package repository

import "time"

const (
	PostTable           = "post_main"
	PostAttachmentTable = "post_attachment"
	PostCategoryTable   = "post_category"
	PostExtendTable     = "post_extend"
	SettingsTable       = "setting"
	UserTable           = "user"
)

type BaseModel struct {
	ID        int64      `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"  gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"  gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"  gorm:"column:deleted_at"`
	Deleted   int        `json:"deleted" form:"deleted"  gorm:"column:deleted"`
}
