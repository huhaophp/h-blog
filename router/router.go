package router

import (
	"bbs/app/admin"
	"bbs/app/middleware"
	"bbs/app/web"
	response "bbs/library"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	authController := new(admin.AuthController)
	homeController := new(admin.HomeController)
	fileController := new(admin.FileController)
	cateController := new(admin.CategoryController)
	articleController := new(admin.ArtcileController)

	// admin routes.
	s.Group("/admin", func(group *ghttp.RouterGroup) {
		group.GET("login", authController.Login)
		group.POST("login", authController.Login)
		group.Middleware(middleware.AdminAuthCheck)
		group.POST("logout", authController.Logout)
		group.GET("home", homeController.Home)
		group.GET("categories", cateController.List)
		group.GET("categories/add", cateController.Add)
		group.POST("categories/add", cateController.Add)
		group.GET("categories/{id}/edit", cateController.Edit)
		group.POST("categories/{id}/edit", cateController.Edit)
		group.POST("categories/{id}/delete", cateController.Delete)
		group.GET("articles", articleController.List)
		group.GET("articles/add", articleController.Add)
		group.POST("articles/add", articleController.Add)
		group.GET("articles/{id}/edit", articleController.Edit)
		group.POST("articles/{id}/edit", articleController.Edit)
		group.POST("articles/{id}/delete", articleController.Delete)
		group.POST("file", fileController.Store)
		group.POST("markdown/file", fileController.MarkdownFileStore)
	})

	// web routes.
	webController := new(web.WebController)
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.InitializeLayoutGlobalViewVariables)
		group.GET("/", webController.Home)
		group.GET("/articles/{id}", webController.Show)
	})

	// Handling 404 pages
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		response.NotFoundView(r)
	})
}
