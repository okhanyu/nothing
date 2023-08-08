package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nothing/config/blog"
	"nothing/internal/app/blog/controller/post"
	"nothing/internal/app/blog/controller/setting"
)

func PostRouteRegister(pc *post.PostController) func(e *gin.Engine) {
	return func(e *gin.Engine) {
		rootUri, postUri := "/", "/post/"
		postUri = fmt.Sprintf("%s%s", blog.Global.System.ApiPrefix, postUri)
		postGroup := e.Group(postUri)
		{
			postGroup.GET("batch/query", pc.FindBatch)
			postGroup.GET(":id/query", pc.FindByID)
			postGroup.GET("partition/batch/query", pc.FindBatchPartition)
		}

		rootGroup := e.Group(rootUri)
		{
			rootGroup.GET("feed", pc.Rss)
			rootGroup.GET("index.xml", pc.Rss)
			rootGroup.GET("blog/index.xml", pc.Rss)
			rootGroup.GET("blog/feed", pc.Rss)
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
