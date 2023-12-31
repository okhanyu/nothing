package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nothing/config"
	postcontroller "nothing/internal/app/controller/post"
	"nothing/internal/app/model/post"
	postservice "nothing/internal/app/service/post"
	settingservice "nothing/internal/app/service/setting"
	pkgrss "nothing/internal/pkg/rss"
	"nothing/pkg/cos"
	"nothing/pkg/response"
)

func init() {
	//variable.Regs = append(variable.Regs, InitSystemApi)
}

type SystemController struct {
	SettingService *settingservice.SettingService
	PostService    *postservice.PostService
	PostAssembler  *postcontroller.PostAssembler
	CosOperate     *cos.Cos
}

func NewSystemController(postServ *postservice.PostService, settingServ *settingservice.SettingService,
	postAssembler *postcontroller.PostAssembler, cosOperate *cos.Cos) *SystemController {
	return &SystemController{
		PostService:    postServ,
		SettingService: settingServ,
		PostAssembler:  postAssembler,
		CosOperate:     cosOperate,
	}
}

func (sc *SystemController) QueryCredential(c *gin.Context) {
	credential := sc.CosOperate.StsCos()
	//if err != nil {
	//	c.JSON(http.StatusOK, response.ErrUnknown)
	//	return
	//}
	c.JSON(http.StatusOK, response.Success.BuildData(credential))
}

func (sc *SystemController) Rss(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}

	if req.Filters.Type == nil || len(req.Filters.Type) == 0 {
		req.Filters.Type = []int{3}
	}
	//if req.Filters.Mode == "" {
	//	req.Filters.Mode = postcons.Simple
	//}
	posts, err := sc.PostService.FindBatch(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	settings, err := sc.SettingService.FindBatch([]int{config.Global.Business.RssSys})
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	if settings == nil || len(settings) == 0 {
		c.JSON(http.StatusOK, response.ErrNotFound)
		return
	}
	if posts == nil || len(posts) == 0 {
		c.JSON(http.StatusOK, response.ErrNotFound)
		return
	}

	rss := pkgrss.Rss(settings[0], sc.PostAssembler.AssemblePostBoToVoForSimpleList(posts))

	//switch strings.ToUpper(rssType.Type) {
	//case "RSS":
	//	mimetype = "application/rss+xml"
	//	response = generateFeeds(data, "RSS")
	//case "ATOM":
	//	mimetype = "application/atom+xml"
	//	response = generateFeeds(data, "ATOM")
	//case "JSON":
	//	mimetype = "application/feed+json"
	//	response = generateFeeds(data, "JSON")
	//}
	c.Data(http.StatusOK, "application/xml", []byte(rss))
}
