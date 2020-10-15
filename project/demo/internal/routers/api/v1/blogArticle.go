package v1

import "github.com/gin-gonic/gin"

type BlogArticle struct {
}

func NewBlogArticle() *BlogArticle {
	return &BlogArticle{}
}

func (blogArticle *BlogArticle) ArticleList(c *gin.Context) {

}

func (blogArticle *BlogArticle) ArticleInfo(c *gin.Context) {

}

func (blogArticle *BlogArticle) CreateArticle(c *gin.Context) {

}

func (blogArticle *BlogArticle) UpdateArticle(c *gin.Context) {

}

func (blogArticle *BlogArticle) DeleteArticle(c *gin.Context) {

}
