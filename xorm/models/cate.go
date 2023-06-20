package models

type Cate struct {
	Aid     int    `json:"aid" xorm:"not null pk autoincr comment('分类id') INT(11)"`
	Cname   string `json:"cname" xorm:"not null default '' comment('分类名称') VARCHAR(200)"`
	Cdesc   string `json:"cdesc" xorm:"not null default '' comment('分类描述') VARCHAR(200)"`
	Isdel   int    `json:"isdel" xorm:"default 0 comment('是否删除{0:正常,1删除}') INT(11)"`
	Addtime int    `json:"addtime" xorm:"not null default 0 comment('添加时间') INT(11)"`
}
