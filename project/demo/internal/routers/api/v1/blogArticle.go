package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type BlogArticle struct {
}

func NewBlogArticle() *BlogArticle {
	return &BlogArticle{}
}

func (blogArticle *BlogArticle) ArticleList(c *gin.Context) {
	//resp.Success(c,"sdfsdf")
	//resp.Error(c,err.NewError(500,"服务器错误"))
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
