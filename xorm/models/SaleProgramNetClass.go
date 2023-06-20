package models

type Saleprogramnetclass struct {
	Rid           int    `json:"rid" xorm:"not null pk autoincr comment('自增') INT(11)"`
	Programid     int    `json:"ProgramId" xorm:"comment('活动ID') INT(11)"`
	Programruleid int    `json:"ProgramRuleId" xorm:"comment('规则ID') INT(11)"`
	Netclassid    int    `json:"NetClassId" xorm:"comment('课程ID') INT(11)"`
	Price         string `json:"Price" xorm:"comment('课程表中的ACTUALPRICE') DECIMAL(8,2)"`
	Actualprice   string `json:"ActualPrice" xorm:"comment('活动价格') DECIMAL(8,2)"`
	Status        int    `json:"status" xorm:"default 1 comment('活动课程是否有效（预留）') INT(11)"`
	Offprice      string `json:"OffPrice" xorm:"comment('直降的金额') DECIMAL(8,2)"`
	Maxnum        int    `json:"MaxNum" xorm:"comment('限购个数') INT(11)"`
	Point         int    `json:"point" xorm:"comment('活动时返的积分') INT(11)"`
	Ratio         string `json:"Ratio" xorm:"default 0.00 comment('折扣 1，2，3') DECIMAL(10,2)"`
	Activebuynum  int    `json:"activeBuyNum" xorm:"default 0 comment('活动开始后课程购买数量') INT(11)"`
}
