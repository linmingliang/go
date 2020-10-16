package v1

import (
	"demo/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type BlogArticle struct {
}

func NewBlogArticle() *BlogArticle {
	return &BlogArticle{}
}

func (blogArticle *BlogArticle) ArticleList(c *gin.Context) {
	logger.InfoF("%s", "ddd")
}

func (blogArticle *BlogArticle) ArticleInfo(c *gin.Context) {
	fmt.Println(c.Keys)
}

func (blogArticle *BlogArticle) CreateArticle(c *gin.Context) {

}

func (blogArticle *BlogArticle) UpdateArticle(c *gin.Context) {

}

func (blogArticle *BlogArticle) DeleteArticle(c *gin.Context) {

}
