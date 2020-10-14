package model

import "time"

type BlogArticle struct {
	Id int32
	//文章标题
	Title string
	//文章描述
	Desc string
	//标签id
	LabelId int32
	//文章内容
	Context string
	//状态0正常，1禁用，2删除
	State int8
	//新建时间
	CreateTime time.Time
	//修改时间
	UpdateTime time.Time
}

func (a *BlogArticle) TableName() string {
	return "blog_article"
}
