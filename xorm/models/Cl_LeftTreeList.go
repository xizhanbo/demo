package models

type ClLefttreelist struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name        string `json:"name" xorm:"not null VARCHAR(255)"`
	Url         string `json:"url" xorm:"VARCHAR(255)"`
	Pid         int    `json:"pid" xorm:"not null default 0 INT(11)"`
	Status      int    `json:"status" xorm:"not null default 1 INT(11)"`
	Css         string `json:"CSS" xorm:"VARCHAR(50)"`
	Isshow      int    `json:"isshow" xorm:"default 1 INT(11)"`
	Showorder   int    `json:"showOrder" xorm:"default 0 comment('显示的顺序') INT(11)"`
	Welcome     string `json:"welcome" xorm:"comment('欢迎页显示 按部门分类，一个节点可以配置多个部门') VARCHAR(50)"`
	Externalurl int    `json:"ExternalURL" xorm:"default 0 comment('是否是外部url，需要拼接参数后边可添加1.拼后台管理员用户名') TINYINT(4)"`
}
