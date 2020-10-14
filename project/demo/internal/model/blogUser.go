package model

import "time"

type BlogUser struct {
	Id int32
	//用户名
	Username string
	//密码
	Password string
	//手机号
	Tel string
	//电子邮箱
	Email string
	//状态0正常1禁用2删除
	State int8
	//新建时间
	CreateTime time.Time
	//修改时间
	UpdateTime time.Time
}

func (b *BlogUser) TableName() string {
	return "blog_user"
}
