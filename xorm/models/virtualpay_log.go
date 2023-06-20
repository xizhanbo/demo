package models

type VirtualpayLog struct {
	Id          int    `json:"id" xorm:"not null pk autoincr comment('ID') INT(11)"`
	OaNumber    string `json:"oa_number" xorm:"not null comment('OA流程单号') VARCHAR(18)"`
	Username    string `json:"username" xorm:"not null comment('用户名') VARCHAR(32)"`
	Realname    string `json:"realname" xorm:"not null comment('真实姓名') VARCHAR(32)"`
	Type        int    `json:"type" xorm:"comment('充值类别') TINYINT(1)"`
	Zhekou      string `json:"zhekou" xorm:"comment('充值时的折扣') DECIMAL(3,2)"`
	Number      string `json:"number" xorm:"comment('充值金额') DECIMAL(10,2)"`
	CreateTime  int    `json:"create_time" xorm:"comment('充值时间') INT(11)"`
	Operater    string `json:"operater" xorm:"comment('操作员') VARCHAR(32)"`
	Comments    string `json:"comments" xorm:"not null comment('充值备注') VARCHAR(255)"`
	Modifytype  int    `json:"ModifyType" xorm:"comment('操作类型 1.充值 ，2修改') INT(11)"`
	Beforemoney string `json:"BeforeMoney" xorm:"comment('修改前的余额') DECIMAL(10)"`
}
