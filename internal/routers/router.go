package routers

import (
	"github.com/FlyInThesky10/TikTok-Fly/global"
	"github.com/FlyInThesky10/TikTok-Fly/internal/controller"
	"github.com/FlyInThesky10/TikTok-Fly/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	regisMiddleWare(r)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		/*apiRouter.GET("/feed/", v1.Feed)*/
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/user/register/", controller.UserRegister)
		apiRouter.POST("/user/login/", controller.UserLogin)
		/*apiRouter.POST("/publish/action/", v1.Publish)
		apiRouter.GET("/publish/list/", v1.PublishList)*/

		// extra apis - I
		/*apiRouter.POST("/favorite/action/", v1.FavoriteAction)
		apiRouter.GET("/favorite/list/", v1.FavoriteList)
		apiRouter.POST("/comment/action/", v1.CommentAction)
		apiRouter.GET("/comment/list/", v1.CommentList)*/

		// extra apis - II
		/*apiRouter.POST("/relation/action/", v1.RelationAction)
		apiRouter.GET("/relation/follow/list/", v1.FollowList)
		apiRouter.GET("/relation/follower/list/", v1.FollowerList)*/
	}

	return r
}

func regisMiddleWare(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
}
