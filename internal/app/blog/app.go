package blog

import (
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"nothing/config/blog"
	"nothing/internal/app/blog/controller/post"
	"nothing/internal/app/blog/controller/setting"
	"nothing/internal/app/blog/controller/system"
	"nothing/internal/app/blog/controller/user"
	postrepo "nothing/internal/app/blog/repository/post"
	settingrepo "nothing/internal/app/blog/repository/setting"
	userrepo "nothing/internal/app/blog/repository/user"
	postservice "nothing/internal/app/blog/service/post"
	settingservice "nothing/internal/app/blog/service/setting"
	userservice "nothing/internal/app/blog/service/user"
	"nothing/internal/pkg/database"
	"nothing/pkg/jwt"
	"nothing/pkg/response"
	"nothing/pkg/server"
	"strings"
)

func CreateBlogServer(name string, conf *blog.Config) (*server.Server, error) {

	dbDsn := conf.System.Db
	db := database.NewDataBase(dbDsn)

	jwt.JwtKey = []byte(conf.System.JwtKey)

	settingService := settingservice.NewSettingService(settingrepo.NewSettingRepository(db))
	postService := postservice.NewPostService(postrepo.NewPostRepository(db))
	postAssembler := post.NewAssembler()

	newServer, err := server.NewServer(
		name,
		server.WithFilter(
			server.CorsMiddleware(),
			func(ctx *gin.Context) {
				if ctx.Request.Method == http.MethodGet || strings.Contains(ctx.Request.RequestURI, "login") {
					ctx.Next()
					return
				}
				authHeader := ctx.GetHeader("Authorization")
				tokenString := ""
				if strings.HasPrefix(authHeader, "Bearer ") {
					// 提取令牌部分
					tokenString = authHeader[7:]
				}
				result, err := jwt.Verify(tokenString)
				if err != nil || result == false {
					ctx.JSON(http.StatusOK, response.ErrJwtToken.BuildData(err.(*jwt2.ValidationError).Error()))
					ctx.Abort()
					return
				}
				ctx.Next()
			},
		),
		server.WithRouter(
			PostRouteRegister(post.NewPostController(postService, postAssembler)),
			SettingRouteRegister(setting.NewSettingController(settingService)),
			SystemRouteRegister(system.NewSystemController(postService, settingService, postAssembler)),
			UserRouteRegister(user.NewUserController(userservice.NewUserService(userrepo.NewUserRepository(db)))),
		),
	)
	if err != nil {
		return nil, err
	}
	return newServer, nil
}
