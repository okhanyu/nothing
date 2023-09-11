package user

import (
	"nothing/internal/app/repository"
)

type User struct {
	repository.BaseModel
	Username string `json:"username" form:"username" gorm:"column:username"`
	Password string `json:"password" form:"password" gorm:"column:password"`
	Role     int    `json:"role" form:"role" gorm:"column:role"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"column:avatar"`
}

func (User) TableName() string {
	return repository.UserTable
}
