package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	_ "github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r.Use(middleware.Translations())
	// 指定swagger的doc.json路径
	// url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	url := ginSwagger.URL("doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	article := v1.NewArticle()
	tag := v1.NewTag()

	// 文件上传路由
	upload := v1.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 文件静态资源访问
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
