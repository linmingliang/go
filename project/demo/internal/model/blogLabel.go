package model

import "time"

type BlogLabel struct {
	Id int32
	//标签名称
	Name string
	//标签类型 1类目2tag
	Type int8
	//状态 0正常1禁用2删除
	State int8
	//新建时间
	CreateTime time.Time
	//修改时间
	UpdateTime time.Time
}

func (b *BlogLabel) TableName() string {
	return "blog_label"
}
