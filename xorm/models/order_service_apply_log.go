package models

type OrderServiceApplyLog struct {
	Id                  int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	OrderId             int    `json:"order_id" xorm:"not null default 0 comment('订单id') index(idx_type_order_id) INT(11)"`
	ClassId             int    `json:"class_id" xorm:"not null default 0 comment('课程id') INT(11)"`
	RedenvelopesPrice   int    `json:"redenvelopes_price" xorm:"not null default 0 comment('红包金额') INT(11)"`
	OtherNotRefundMoney int    `json:"other_not_refund_money" xorm:"not null default 0 comment('图书、课时及申论批改扣费') INT(11)"`
	RealAmount          int    `json:"real_amount" xorm:"not null default 0 comment('实退金额') INT(11)"`
	ReceiptMoney        int    `json:"receipt_money" xorm:"not null comment('原班价格') INT(11)"`
	RefundType          int    `json:"refund_type" xorm:"not null default 0 comment('退去方式') TINYINT(1)"`
	RefundWhereabouts   int    `json:"refund_whereabouts" xorm:"not null default 0 comment('退款去向') TINYINT(1)"`
	Reason              int    `json:"reason" xorm:"not null default 0 comment('原因') TINYINT(1)"`
	NewClassPrice       int    `json:"new_class_price" xorm:"not null default 0 comment('新班价钱') INT(11)"`
	NewClassId          int    `json:"new_class_id" xorm:"not null default 0 comment('新班价钱') INT(11)"`
	Region              string `json:"region" xorm:"not null comment('省市区 逗号分隔') VARCHAR(255)"`
	UserAddress         string `json:"user_address" xorm:"not null default '' VARCHAR(255)"`
	Phone               string `json:"phone" xorm:"not null default '' comment('用户电话') VARCHAR(11)"`
	ClassDiffPrice      int    `json:"class_diff_price" xorm:"not null default 0 comment('0') INT(11)"`
	Remark              string `json:"remark" xorm:"not null comment('备注') TEXT"`
	Operator            string `json:"operator" xorm:"not null VARCHAR(100)"`
	Status              int    `json:"status" xorm:"not null default 0 TINYINT(4)"`
	UpdateTime          int    `json:"update_time" xorm:"not null default 0 INT(11)"`
	CreateTime          int    `json:"create_time" xorm:"not null default 0 INT(11)"`
	Type                string `json:"type" xorm:"not null default '' comment('类型') index(idx_type_order_id) VARCHAR(20)"`
}
