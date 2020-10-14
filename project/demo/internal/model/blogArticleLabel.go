package model

type BlogArticleLabel struct {
	Id        int32
	ArticleId int32
	LabelId   int32
}

func (b *BlogArticleLabel) TableName() string {
	return "blog_article_label"
}
