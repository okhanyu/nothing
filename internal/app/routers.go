package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nothing/config"
	"nothing/internal/app/controller/post"
	"nothing/internal/app/controller/setting"
	"nothing/internal/app/controller/system"
	"nothing/internal/app/controller/user"
)

func UserRouteRegister(user *user.UserController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		rootUri := "/"
		postGroup := e.Group(rootUri)
		{
			postGroup.POST("login", user.Login)
		}
	}
}

func PostRouteRegister(pc *post.PostController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		postUri := "/post/"
		postUri = fmt.Sprintf("%s%s", config.Global.System.ApiPrefix, postUri)
		postGroup := e.Group(postUri)
		{
			postGroup.POST("batch/query", pc.FindBatch)
			postGroup.GET("query", pc.FindByID)
			postGroup.POST("partition/batch/query", pc.FindBatchPartition)
			postGroup.POST("attachment/content/query", pc.FindAttachmentContent)
			postGroup.POST("create", pc.CreatePost)
		}
	}

}

func SettingRouteRegister(sc *setting.SettingController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		_, settingUri := "/", "/setting/"
		settingUri = fmt.Sprintf("%s%s", config.Global.System.ApiPrefix, settingUri)
		settingGroup := e.Group(settingUri)
		{
			settingGroup.POST("batch/query", sc.QuerySetting)
			//settingGroup.POST("setting/:id/query", deleteProduct)
		}
	}
}

func SystemRouteRegister(sc *system.SystemController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		rootUri, systemUri := "/", "/system/"
		systemUri = fmt.Sprintf("%s%s", config.Global.System.ApiPrefix, systemUri)
		systemGroup := e.Group(systemUri)
		{
			systemGroup.POST("sts/query", sc.QueryCredential)
			//settingGroup.POST("setting/:id/query", deleteProduct)
		}

		rootGroup := e.Group(rootUri)
		{
			rootGroup.GET("feed", sc.Rss)
			rootGroup.GET("index.xml", sc.Rss)
			rootGroup.GET("blog/index.xml", sc.Rss)
			rootGroup.GET("blog/feed", sc.Rss)
		}

	}
}
