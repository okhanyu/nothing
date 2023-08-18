package system

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nothing/config/blog"
	postcons "nothing/internal/app/blog/cons/post"
	postcontroller "nothing/internal/app/blog/controller/post"
	"nothing/internal/app/blog/model/post"
	postservice "nothing/internal/app/blog/service/post"
	settingservice "nothing/internal/app/blog/service/setting"
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
}

func NewSystemController(postServ *postservice.PostService, settingServ *settingservice.SettingService,
	postAssembler *postcontroller.PostAssembler) *SystemController {
	return &SystemController{
		PostService:    postServ,
		SettingService: settingServ,
		PostAssembler:  postAssembler,
	}
}

func (sc *SystemController) QueryCredential(c *gin.Context) {
	credential := cos.StsCos()
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

	if req.Type == nil || len(req.Type) == 0 {
		req.Type = []int{3}
	}
	if req.Mode == "" {
		req.Mode = postcons.Simple
	}
	posts, err := sc.PostService.FindBatch(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	settings, err := sc.SettingService.FindBatch([]int{blog.Global.Business.RssSys})
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
