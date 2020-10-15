package v1

import "github.com/gin-gonic/gin"

type BlogUser struct {
}

func NewBlogUser() *BlogUser {
	return &BlogUser{}
}
func (blogUser *BlogUser) BlogUserList(c *gin.Context) {

}
func (blogUser *BlogUser) BlogUserInfo(c *gin.Context) {

}
func (blogUser *BlogUser) CreateBlogUser(c *gin.Context) {

}
func (blogUser *BlogUser) UpdateBlogUser(c *gin.Context) {

}
func (blogUser *BlogUser) DeleteBlogUser(c *gin.Context) {

}
