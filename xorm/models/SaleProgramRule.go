package models

import (
	"time"
)

type Saleprogramrule struct {
	Rid              int       `json:"rid" xorm:"not null pk autoincr comment('自增') INT(11)"`
	Programid        int       `json:"ProgramId" xorm:"comment('活动ID') INT(11)"`
	Ruletitle        string    `json:"RuleTitle" xorm:"comment('规则名') VARCHAR(100)"`
	Fullmoney        string    `json:"FullMoney" xorm:"comment('满N钱') DECIMAL(8,2)"`
	Offmoney         string    `json:"OffMoney" xorm:"comment('减N钱') DECIMAL(10)"`
	Vouchernum       int       `json:"voucherNum" xorm:"comment('代金券个数') INT(11)"`
	Vouchercode      string    `json:"voucherCode" xorm:"comment('代金券编号') VARCHAR(50)"`
	Vouchermoney     string    `json:"voucherMoney" xorm:"comment('代金券金额') DECIMAL(8,2)"`
	Voucherenddate   time.Time `json:"voucherEndDate" xorm:"comment('代金券到期时间') DATETIME"`
	Groupmoney       string    `json:"groupMoney" xorm:"comment('组合的总金额') DECIMAL(8,2)"`
	Voucherfullmoney string    `json:"voucherFullMoney" xorm:"comment('代金券满减') DECIMAL(8,2)"`
}
