package models

import (
	"time"
)

type Saleactivty struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Description string    `json:"description" xorm:"comment('描述') VARCHAR(500)"`
	Rid         int       `json:"rid" xorm:"not null comment('代金券id') INT(11)"`
	Tid         int       `json:"tid" xorm:"not null comment('标题id') INT(11)"`
	Createdate  time.Time `json:"createDate" xorm:"comment('创建时间') DATETIME"`
	Editdate    time.Time `json:"editDate" xorm:"comment('编辑时间') DATETIME"`
}
