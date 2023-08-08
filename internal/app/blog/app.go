package blog

import (
	"nothing/config/blog"
	"nothing/internal/app/blog/controller/post"
	"nothing/internal/app/blog/controller/setting"
	postrepo "nothing/internal/app/blog/repository/post"
	settingrepo "nothing/internal/app/blog/repository/setting"
	postservice "nothing/internal/app/blog/service/post"
	settingservice "nothing/internal/app/blog/service/setting"
	"nothing/internal/pkg/database"
	"nothing/pkg/server"
)

func CreateBlogServer(name string, conf *blog.Config) (*server.Server, error) {

	dbDsn := conf.System.Db

	db := database.NewDataBase(dbDsn)

	postRepository := postrepo.NewPostRepository(db)
	postService := postservice.NewPostService(postRepository)

	settingRepository := settingrepo.NewSettingRepository(db)
	settingService := settingservice.NewSettingService(settingRepository)

	postController := post.NewPostController(postService, settingService)
	settingController := setting.NewSettingController(settingService)

	newServer, err := server.NewServer(
		name,
		server.WithFilter(),
		server.WithRouter(
			PostRouteRegister(postController),
			SettingRouteRegister(settingController),
		),
	)
	if err != nil {
		return nil, err
	}
	return newServer, nil
}
