package models

type Zengsaleprogramnetclass struct {
	Rid         int    `json:"rid" xorm:"not null pk autoincr comment('自增') INT(11)"`
	Programid   int    `json:"ProgramId" xorm:"comment('活动ID') INT(11)"`
	Netclassid  int    `json:"NetClassId" xorm:"comment('课程ID') INT(11)"`
	Price       string `json:"Price" xorm:"comment('课程表中的ACTUALPRICE') DECIMAL(8,2)"`
	Actualprice string `json:"ActualPrice" xorm:"comment('活动价格') DECIMAL(8,2)"`
	Title       string `json:"Title" xorm:"not null comment('课程标题') VARCHAR(200)"`
	Iszengsong  int    `json:"IsZengsong" xorm:"not null comment('是否为赠送') INT(11)"`
}
