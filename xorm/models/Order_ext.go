package models

type OrderExt struct {
	Id           int64  `json:"id" xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	Orderid      int64  `json:"orderId" xorm:"not null comment('订单id') BIGINT(20)"`
	Userid       int64  `json:"userId" xorm:"comment('用户id') BIGINT(20)"`
	Province     string `json:"province" xorm:"default '' comment('所属省份') VARCHAR(50)"`
	City         string `json:"city" xorm:"default '' comment('所属城市') VARCHAR(50)"`
	Exampost     string `json:"examPost" xorm:"default '' comment('报考职位') VARCHAR(100)"`
	Exampostcode string `json:"examPostCode" xorm:"default '' comment('职位代码') VARCHAR(50)"`
	Tjcode       string `json:"tjCode" xorm:"default '' comment('站内统计代码') VARCHAR(100)"`
	Source       string `json:"source" xorm:"default '' comment('来源，部门订单（mk：市场，rec：猜你喜欢，zx：华图在线），ztk：砖题库') VARCHAR(20)"`
	Platform     string `json:"platform" xorm:"default '' comment('使用平台（PC端：pc，M站：m，安卓客户端：Android），苹果客户端：ios') VARCHAR(20)"`
	Area         string `json:"area" xorm:"default '' comment('所属区') VARCHAR(50)"`
	Consignee    string `json:"conSignee" xorm:"default '' comment('收货人姓名') VARCHAR(50)"`
	Address      string `json:"address" xorm:"default '' comment('收货地址') VARCHAR(255)"`
	Zipcode      string `json:"zipCode" xorm:"comment('邮编') VARCHAR(50)"`
	Phone        string `json:"phone" xorm:"default '' comment('手机号') CHAR(11)"`
	Email        string `json:"email" xorm:"default '' comment('邮箱') VARCHAR(50)"`
	Delivertype  int    `json:"deliverType" xorm:"comment('邮寄方式') INT(11)"`
	Needinvoice  int    `json:"needInvoice" xorm:"comment('不为空（是否需要发票）') TINYINT(3)"`
	Invoicehead  string `json:"invoiceHead" xorm:"default '' comment('发票抬头') VARCHAR(255)"`
	Remark       string `json:"remark" xorm:"comment('订单备注') VARCHAR(255)"`
	Operremark   string `json:"operRemark" xorm:"default '' comment('操作人员备注') VARCHAR(255)"`
	Invoiced     int    `json:"invoiced" xorm:"default 0 comment('是否已开发票') TINYINT(3)"`
	Operuser     string `json:"operUser" xorm:"comment('操作人') VARCHAR(50)"`
	Receipted    int    `json:"receipted" xorm:"comment('是否开发收据，0：没 1：开') TINYINT(3)"`
	Otherremark  string `json:"otherRemark" xorm:"comment('其他备注') TEXT"`
}
