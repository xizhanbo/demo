package models

import (
	"time"
)

type Netclassesqqgroup struct {
	Rid             int       `json:"rid" xorm:"not null pk INT(11)"`
	Weburl          string    `json:"WebUrl" xorm:"comment('QQ一键加群网页链接') VARCHAR(255)"`
	Androidfunction string    `json:"AndroidFunction" xorm:"comment('安卓一键加群') TEXT"`
	Iosfunction     string    `json:"IosFunction" xorm:"comment('Ios一键加群') TEXT"`
	Statusdate      time.Time `json:"StatusDate" xorm:"not null comment('时间') DATETIME"`
	Qrcode          string    `json:"QrCode" xorm:"comment('QQ二维码') VARCHAR(255)"`
	Qqnum           int       `json:"QQnum" xorm:"not null comment('QQ号码') INT(11)"`
	Wechatnumber    string    `json:"wechatNumber" xorm:"comment('微信号') VARCHAR(200)"`
	Wechatqrcode    string    `json:"wechatQrCode" xorm:"comment('微信二维码') CHAR(255)"`
	Service         int       `json:"service" xorm:"default 0 comment('服务渠道1QQ2微信3小程序（关注公众号）4小程序（联系教师）5加群计划') TINYINT(1)"`
	Accountqrcode   string    `json:"accountQrCode" xorm:"comment('公众号二维码') VARCHAR(255)"`
	Firstsubtitle   string    `json:"firstSubtitle" xorm:"comment('副标题一') VARCHAR(200)"`
	Secondsubtitle  string    `json:"secondSubtitle" xorm:"comment('副标题二') VARCHAR(200)"`
	Planqrcode      string    `json:"PlanQrCode" xorm:"comment('计划加群二维码') VARCHAR(255)"`
}
