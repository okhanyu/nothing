package setting

import (
	"github.com/gin-gonic/gin"
	"net/http"
	settingservice "nothing/internal/app/service/setting"
	"nothing/pkg/response"
)

func init() {
	//variable.Regs = append(variable.Regs, InitSettingApi)
}

type SettingController struct {
	Service *settingservice.SettingService
}

func NewSettingController(service *settingservice.SettingService) *SettingController {
	return &SettingController{
		Service: service,
	}
}

func (sc *SettingController) QuerySetting(c *gin.Context) {
	//ids := c.QueryArray("sys")
	var idsInt []int
	err := c.ShouldBindJSON(&idsInt)
	if err != nil {
		return
	}
	//var idsInt []int
	//for _, id := range ids {
	//	i, err := strconv.Atoi(id)
	//	if err != nil {
	//		continue
	//	}
	//	idsInt = append(idsInt, i)
	//}
	list, err := sc.Service.FindBatch(idsInt)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(list))
}
