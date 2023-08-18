package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nothing/config/blog"
	"nothing/internal/app/blog/controller/post"
	"nothing/internal/app/blog/controller/setting"
	"nothing/internal/app/blog/controller/system"
	"nothing/internal/app/blog/controller/user"
)

func UserRouteRegister(user *user.UserController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		rootUri := "/"
		//rootUri = fmt.Sprintf("%s%s", blog.Global.System.ApiPrefix, rootUri)
		postGroup := e.Group(rootUri)
		{
			postGroup.POST("login", user.Login)
		}
	}
}

func PostRouteRegister(pc *post.PostController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		postUri := "/post/"
		postUri = fmt.Sprintf("%s%s", blog.Global.System.ApiPrefix, postUri)
		postGroup := e.Group(postUri)
		{
			postGroup.GET("batch/query", pc.FindBatch)
			postGroup.GET(":id/query", pc.FindByID)
			postGroup.GET("partition/batch/query", pc.FindBatchPartition)
			postGroup.POST("create", pc.CreatePost)
			postGroup.GET("attachment/content/query", pc.FindAttachmentContent)
		}
	}

}

func SettingRouteRegister(sc *setting.SettingController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		_, settingUri := "/", "/setting/"
		settingUri = fmt.Sprintf("%s%s", blog.Global.System.ApiPrefix, settingUri)
		settingGroup := e.Group(settingUri)
		{
			settingGroup.GET("batch/query", sc.QuerySetting)
			//settingGroup.POST("setting/:id/query", deleteProduct)
		}
	}
}

func SystemRouteRegister(sc *system.SystemController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		rootUri, systemUri := "/", "/system/"
		systemUri = fmt.Sprintf("%s%s", blog.Global.System.ApiPrefix, systemUri)
		systemGroup := e.Group(systemUri)
		{
			systemGroup.GET("sts/query", sc.QueryCredential)
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
