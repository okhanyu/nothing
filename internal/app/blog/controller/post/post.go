package post

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	postcons "nothing/internal/app/blog/cons/post"
	"nothing/internal/app/blog/model/post"
	postservice "nothing/internal/app/blog/service/post"
	"nothing/pkg/response"
	"strconv"
)

func init() {
	//variable.Regs = append(variable.Regs, InitPostApi)
}

type PostController struct {
	Service       *postservice.PostService
	PostAssembler *PostAssembler
}

func NewPostController(serv *postservice.PostService, postAssembler *PostAssembler) *PostController {
	return &PostController{
		Service:       serv,
		PostAssembler: postAssembler,
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
	c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePartitionPostBoToVo(list)))
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
		c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForHiddenList(list)))
		break
	case postcons.Simple:
		c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForSimpleList(list)))
		break
	case postcons.Normal:
		c.JSON(http.StatusOK, response.Success.BuildData(pc.PostAssembler.AssemblePostBoToVoForNormalList(list, req.RespAttachmentMode)))
		break
	default:
		c.JSON(http.StatusOK, response.ErrMissingParam)
	}
	return
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

func (pc *PostController) FindAttachmentContent(c *gin.Context) {
	var req post.FindReq
	err := c.ShouldBindQuery(&req)
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
