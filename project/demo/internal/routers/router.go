package routers

import (
	"demo/internal/middleware"
	v1 "demo/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery(), middleware.BeforeRequest())

	r := engine.Group("/api/v1")
	{
		article := v1.NewBlogArticle()
		r.GET("/articles", article.ArticleList)
		r.GET("/articles/:id", article.ArticleInfo)
		r.POST("/articles", article.CreateArticle)
		r.PUT("/articles/;id", article.UpdateArticle)
		r.PATCH("/articles/:id/:state", article.UpdateArticle)
		r.DELETE("/articles/:id", article.DeleteArticle)

		label := v1.NewBlogLabel()
		r.GET("/labels", label.BlogLabelList)
		r.GET("/labels/:id", label.BlogLabelInfo)
		r.POST("/labels", label.CreateLabel)
		r.PUT("/labels/:id", label.UpdateLabel)
		r.PATCH("/labels/:id/:state", label.UpdateLabel)
		r.DELETE("/labels/:id", label.DeleteLabel)
	}
	return engine
}
