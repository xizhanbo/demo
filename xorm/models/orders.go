package models

import (
	"time"
)

type Orders struct {
	Id                   int64     `json:"id" xorm:"pk autoincr comment('订单ID') BIGINT(20)"`
	Ordernum             string    `json:"orderNum" xorm:"comment('订单编号') VARCHAR(100)"`
	Userid               int64     `json:"userId" xorm:"comment('用户ID') BIGINT(20)"`
	Username             string    `json:"userName" xorm:"comment('用户名') VARCHAR(100)"`
	Paymenttype          int       `json:"paymentType" xorm:"comment('支付方式') INT(11)"`
	Moneyreceipt         int       `json:"moneyReceipt" xorm:"default 0 comment('收到的金额, 分') INT(11)"`
	Moneyfree            int       `json:"moneyFree" xorm:"default 0 comment('优惠金额, 分') INT(11)"`
	Moneyfree2           int       `json:"moneyFree2" xorm:"default 0 comment('促销优惠码优惠, 分') INT(11)"`
	Moneystudycard       int       `json:"moneyStudyCard" xorm:"default 0 comment('学习卡金额, 分') INT(11)"`
	Moneysum             int       `json:"moneySum" xorm:"default 0 comment('付款顶大金额, 分') INT(11)"`
	Moneytotal           int       `json:"moneyTotal" xorm:"default 0 comment('付款订单总额, 分') INT(11)"`
	Discount             int       `json:"discount" xorm:"default 0 comment('折扣金额, 分') INT(11)"`
	Presentexp           int       `json:"presentExp" xorm:"INT(11)"`
	Status               int       `json:"status" xorm:"default 0 comment('订单状态') INT(11)"`
	Begintime            time.Time `json:"beginTime" xorm:"comment('订单课程开始时间') DATETIME"`
	Endtime              time.Time `json:"endTime" xorm:"comment('订单课程失效时间') DATETIME"`
	Addtime              time.Time `json:"addTime" xorm:"comment('下单时间') DATETIME"`
	Paytime              time.Time `json:"payTime" xorm:"comment('支付时间') DATETIME"`
	Domainname           string    `json:"domainName" xorm:"comment('订单来源描述') VARCHAR(150)"`
	Userregdomain        string    `json:"userRegDomain" xorm:"comment('用户注册来源') VARCHAR(150)"`
	Fromadid             int64     `json:"fromAdId" xorm:"comment('联盟ID') BIGINT(20)"`
	Fromuser             string    `json:"fromUser" xorm:"comment('联盟服务地址') VARCHAR(150)"`
	Fromuseradid         int64     `json:"fromUserAdId" xorm:"comment('联盟用户ID') BIGINT(20)"`
	Fromadhost           string    `json:"fromAdHost" xorm:"comment('联盟服务地址') VARCHAR(150)"`
	Classids             string    `json:"classIds" xorm:"comment('课程ID,用逗号隔开') VARCHAR(255)"`
	Orderorigin          int       `json:"orderOrigin" xorm:"comment('订单来源标识') INT(11)"`
	Callbackuserrealname string    `json:"callbackUserRealName" xorm:"comment('外呼员') VARCHAR(100)"`
	Callbacktime         time.Time `json:"callbackTime" xorm:"comment('外呼时间') DATETIME"`
	Callbackstatus       int       `json:"callbackStatus" xorm:"comment('外呼状态') INT(11)"`
	Cashaccount          int       `json:"cashAccount" xorm:"comment('支付订单是使用的现金账户金额, 单位分') INT(11)"`
	Mainclassid          int64     `json:"mainClassId" xorm:"comment('买赠活动中, 主课程ID') BIGINT(20)"`
	Mainorderid          int64     `json:"mainOrderId" xorm:"comment('买赠活动中, 主订单ID') BIGINT(20)"`
	Rechargecost         int       `json:"rechargeCost" xorm:"comment('充值金币消费') INT(11)"`
	Taskchargecost       int       `json:"taskChargeCost" xorm:"comment('任务金币消费') INT(11)"`
	Ishidden             int       `json:"isHidden" xorm:"default 0 comment('是否隐藏,0-否, 1-是') TINYINT(1)"`
}
