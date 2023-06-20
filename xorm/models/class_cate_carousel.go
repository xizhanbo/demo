package models

import (
	"time"
)

type ClassCateCarousel struct {
	Id        int       `json:"id" xorm:"not null pk INT(11)"`
	CateId    int       `json:"cate_id" xorm:"not null comment('分类id') INT(11)"`
	ImgUrl    string    `json:"img_url" xorm:"comment('图片链接') VARCHAR(200)"`
	ActiveUrl string    `json:"active_url" xorm:"comment('活动链接') VARCHAR(200)"`
	Decribe   string    `json:"decribe" xorm:"comment('描述') VARCHAR(255)"`
	Sort      int       `json:"sort" xorm:"comment('排序') INT(10)"`
	Status    int       `json:"status" xorm:"comment('状态 0待审核1上线2下线3删除') TINYINT(1)"`
	CreateAt  time.Time `json:"create_at" xorm:"comment('创建时间') DATETIME"`
	UpdateAt  time.Time `json:"update_at" xorm:"comment('修改时间') DATETIME"`
}
