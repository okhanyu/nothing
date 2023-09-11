package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nothing/internal/app/model/user"
	userservice "nothing/internal/app/service/user"
	"nothing/pkg/response"
)

func init() {
	//variable.Regs = append(variable.Regs, InitUserApi)
}

type UserController struct {
	Service *userservice.UserService
}

func NewUserController(service *userservice.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

func (sc *UserController) Login(c *gin.Context) {
	var param user.UserReq
	err := c.ShouldBindJSON(&param)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}

	userInfo, err := sc.Service.FindByUsernameAndPassword(param)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown.BuildData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(userInfo))
}
