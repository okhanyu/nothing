package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	postcons "nothing/internal/app/blog/cons/post"
	"nothing/internal/app/blog/model/post"
	postservice "nothing/internal/app/blog/service/post"
	settingservice "nothing/internal/app/blog/service/setting"
	pkgrss "nothing/internal/pkg/rss"
	"nothing/pkg/response"
	"strconv"
)

func init() {
	//variable.Regs = append(variable.Regs, InitPostApi)
}

type PostController struct {
	Service        *postservice.PostService
	SettingService *settingservice.SettingService
}

func NewPostController(serv *postservice.PostService, settingServ *settingservice.SettingService) *PostController {
	return &PostController{
		Service:        serv,
		SettingService: settingServ,
	}
}

func (pc *PostController) FindBatchPartition(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	list, err := pc.Service.FindBatchPartition(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(AssemblePartitionPostBoToVo(list)))
}

func (pc *PostController) FindBatch(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	list, err := pc.Service.FindBatch(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	switch req.Mode {
	case postcons.Hidden:
		c.JSON(http.StatusOK, response.Success.BuildData(AssemblePostBoToVoForHiddenList(list)))
		break
	case postcons.Simple:
		c.JSON(http.StatusOK, response.Success.BuildData(AssemblePostBoToVoForSimpleList(list)))
		break
	case postcons.Normal:
		c.JSON(http.StatusOK, response.Success.BuildData(AssemblePostBoToVoForNormalList(list, req.RespAttachmentMode)))
		break
	default:
		c.JSON(http.StatusOK, response.ErrMissingParam)
	}
	return
}

func (pc *PostController) Rss(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindQuery(req)
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
	posts, err := pc.Service.FindBatch(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	settings, err := pc.SettingService.FindBatch([]int{1})
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

	rss := pkgrss.Rss(settings[0], AssemblePostBoToVoForSimpleList(posts))

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

func (pc *PostController) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	postDetail, err := pc.Service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(postDetail))
}
