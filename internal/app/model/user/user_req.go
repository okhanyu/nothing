package user

import (
	"nothing/internal/app/repository"
)

type UserReq struct {
	ID       int64  `json:"id" form:"id"  gorm:"column:id;primaryKey"`
	Username string `json:"username" form:"username" gorm:"column:username"`
	Password string `json:"password" form:"password" gorm:"column:password"`
	Token    string `json:"token" form:"token" gorm:"-"`
}

func (UserReq) TableName() string {
	return repository.UserTable
}
