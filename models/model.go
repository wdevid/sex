package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"

)

var (
	globalnum int
)

func init() {
	globalnum = 100000000
}

type User struct {
	UserName string
	PassWord string `orm:"size(500)"`
	Id       int64 `auto:"true" index:"pk"`
	CId int64
	HeadPortrait string
	IsHead int64 //是否设置头像0是,1否
	HeadPath string
	Session string
}

type Customer struct {
	Id              int64 `auto:"true" index:"pk"`
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	Public          int64
	Uname  string
}

func RegisterDB() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册 model
	orm.RegisterModel(new(User), new(Mimi), new(Customer))
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/myapp?charset=utf8&loc=Asia%2FShanghai", 30, 30) //密码为空格式
}

