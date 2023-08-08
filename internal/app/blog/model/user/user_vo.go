package user

import (
	"time"
)

type UserVo struct {
	ID        int64      `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"  gorm:"column:created_at"`
	Username  string     `json:"username" form:"username" gorm:"column:username"`
	Role      int        `json:"role" form:"role" gorm:"column:role"`
	Avatar    string     `json:"avatar" form:"avatar" gorm:"column:avatar"`
}
