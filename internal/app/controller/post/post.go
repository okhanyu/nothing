package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nothing/internal/app/model/post"
	pkgPostService "nothing/internal/app/service/post"
	"nothing/pkg/response"
)

func init() {
	//variable.Regs = append(variable.Regs, InitPostApi)
}

type PostController struct {
	Service       *pkgPostService.PostService
	PostAssembler *PostAssembler
}

func NewPostController(serv *pkgPostService.PostService, postAssembler *PostAssembler) *PostController {
	return &PostController{
		Service:       serv,
		PostAssembler: postAssembler,
	}
}

// FindBatchPartition 按照分类分组获取标题列表
func (pc *PostController) FindBatchPartition(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	list, err := pc.Service.FindBatchPartition(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePartitionPostBoToVo(list)))
}

// FindBatch 根据条件获取列表
func (pc *PostController) FindBatch(c *gin.Context) {
	var req post.FindReq
	//err := c.ShouldBindQuery(&req)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	list, err := pc.Service.FindBatch(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForNormalList(list, req.RespAttachmentMode)))
	//switch req.Mode {
	//case pkgPostCons.Hidden:
	//	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForHiddenList(list)))
	//	break
	//case pkgPostCons.Simple:
	//	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForSimpleList(list)))
	//	break
	//case pkgPostCons.Normal:
	//	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForNormalList(list, req.RespAttachmentMode)))
	//	break
	//default:
	//	c.JSON(http.StatusOK, response.ErrMissingParam)
	//}
	return
}

// FindByID 根据ID获取详情
func (pc *PostController) FindByID(c *gin.Context) {
	type Param struct {
		Id int64 `json:"id" form:"id"`
	}
	var param Param
	err := c.ShouldBindQuery(&param)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	//id, err := strconv.ParseInt(param.Id, 10, 64)
	//if err != nil {
	//	log.Printf("param is error : %+v ", err)
	//	return
	//}
	postDetail, err := pc.Service.FindByID(param.Id)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(postDetail))
}

// FindAttachmentContent 随机或者按照创建时间排序获取纯附件内容列表
func (pc *PostController) FindAttachmentContent(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	list, err := pc.Service.FindAttachmentContent(req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown)
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(list))
	return
}

// CreatePost 创建
func (pc *PostController) CreatePost(c *gin.Context) {
	var req post.CreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("param is error : %+v ", err)
		return
	}
	err = pc.Service.CreatePost(c, req)
	if err != nil {
		c.JSON(http.StatusOK, response.ErrUnknown.BuildData(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success.BuildData(""))
}
