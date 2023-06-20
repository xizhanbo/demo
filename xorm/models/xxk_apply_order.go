package models

import (
	"time"
)

type XxkApplyOrder struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ApplyOrder   string    `json:"apply_order" xorm:"not null default '0' comment('学习卡申请单编号') CHAR(20)"`
	NumCard10    int       `json:"num_card_10" xorm:"not null default 0 comment('面值10元的学习卡数量') INT(11)"`
	NumCard50    int       `json:"num_card_50" xorm:"not null default 0 comment('面值50元的学习卡数量') INT(11)"`
	NumCard100   int       `json:"num_card_100" xorm:"not null default 0 comment('面值100元的学习卡数量') INT(11)"`
	NumCard200   int       `json:"num_card_200" xorm:"not null default 0 comment('面值200元的学习卡数量') INT(11)"`
	NumCard500   int       `json:"num_card_500" xorm:"not null default 0 comment('面值500元的学习卡数量') INT(11)"`
	NumCardYueka int       `json:"num_card_yueka" xorm:"not null default 0 comment('月卡') INT(11)"`
	NumCardJika  int       `json:"num_card_jika" xorm:"not null default 0 comment('季卡') INT(11)"`
	UploadImg    string    `json:"upload_img" xorm:"not null default '' comment('用户上传的汇款凭证') VARCHAR(100)"`
	Username     string    `json:"username" xorm:"not null default '' comment('申请学习卡的用户名（代理商的用户名）') VARCHAR(20)"`
	Consignee    string    `json:"consignee" xorm:"not null default '' comment('收件人姓名') VARCHAR(20)"`
	Mobilephone  string    `json:"mobilephone" xorm:"not null default '' comment('收件人手机号码') CHAR(11)"`
	Postcode     string    `json:"postcode" xorm:"not null default '' comment('收件人邮政编码') CHAR(6)"`
	Mailaddress  string    `json:"mailaddress" xorm:"not null default '' comment('收件地址') VARCHAR(50)"`
	Message      string    `json:"message" xorm:"not null comment('留言') TEXT"`
	ApplyCardNum int       `json:"apply_card_num" xorm:"not null default 0 comment('申请学习卡的张数') INT(11)"`
	ApplyCardPay int       `json:"apply_card_pay" xorm:"not null default 0 comment('付款金额') INT(11)"`
	ApplyStatus  int       `json:"apply_status" xorm:"not null default 0 comment('申请单状态（未提交：0，待审核：1，已批准：2，已退回：3，已派发：4）') TINYINT(4)"`
	ApplyRemark  string    `json:"apply_remark" xorm:"not null default '' comment('备注') VARCHAR(50)"`
	ApplyTime    time.Time `json:"apply_time" xorm:"not null default CURRENT_TIMESTAMP comment('申请时间') TIMESTAMP"`
}
