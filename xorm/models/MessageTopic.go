package models

import (
	"time"
)

type Messagetopic struct {
	Tid            int       `json:"tid" xorm:"not null pk autoincr comment('短信主题自增id') INT(11)"`
	Department     int       `json:"department" xorm:"comment('1:产品中心2:产品实施3:销售中心4:市场运营中心5:教研6:网站组7:砖题库8:分校') TINYINT(4)"`
	Topic          string    `json:"topic" xorm:"comment('短信主题名') VARCHAR(255)"`
	Messagetemplet string    `json:"messagetemplet" xorm:"comment('短信模板内容') VARCHAR(255)"`
	Createtime     time.Time `json:"createtime" xorm:"comment('创建时间') DATETIME"`
	Createuser     string    `json:"createuser" xorm:"comment('创建人') VARCHAR(255)"`
}
