package app

import (
	"nothing/config"
	pkgPostCtrl "nothing/internal/app/controller/post"
	"nothing/internal/app/controller/setting"
	"nothing/internal/app/controller/system"
	"nothing/internal/app/controller/user"
	pkgPostRepo "nothing/internal/app/repository/post"
	pkgSettingRepo "nothing/internal/app/repository/setting"
	pkgUserRepo "nothing/internal/app/repository/user"
	pkgPostService "nothing/internal/app/service/post"
	pkgSettingService "nothing/internal/app/service/setting"
	pkgUserService "nothing/internal/app/service/user"
	"nothing/internal/pkg/database"
	"nothing/pkg/cos"
	"nothing/pkg/jwt"
	"nothing/pkg/server"
)

func CreateServer(serverName string, conf *config.Config) (*server.Server, error) {

	// 链接数据库
	dbDsn := conf.System.Db
	dbInstance := database.NewDataBase(dbDsn)

	jwt.JwtKey = []byte(conf.System.JwtKey)

	settingRepoInstance := pkgSettingRepo.NewSettingRepository(dbInstance)
	settingService := pkgSettingService.NewSettingService(settingRepoInstance)

	postRepoInstance := pkgPostRepo.NewPostRepository(dbInstance)
	postService := pkgPostService.NewPostService(postRepoInstance)

	postCtrlAssembler := pkgPostCtrl.NewAssembler()

	newServer, err := server.NewServer(
		serverName,
		server.WithFilter(
			server.CorsMiddlewareFilter(),
			server.AuthHandleFilter(),
		),
		server.WithRouter(
			PostRouteRegister(
				pkgPostCtrl.NewPostController(
					postService,
					postCtrlAssembler,
				),
			),
			SettingRouteRegister(
				setting.NewSettingController(
					settingService,
				),
			),
			SystemRouteRegister(
				system.NewSystemController(
					postService,
					settingService,
					postCtrlAssembler,
					cos.NewCos(
						conf.System.CosKey,
						conf.System.CosId,
						conf.System.CosAppid,
						conf.System.CosBucket,
						conf.System.CosRegion,
					),
				),
			),
			UserRouteRegister(
				user.NewUserController(
					pkgUserService.NewUserService(
						pkgUserRepo.NewUserRepository(
							dbInstance,
						),
					),
				),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	return newServer, nil
}
