package models

import (
	"time"
)

type FaceOrder struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName          string    `json:"user_name" xorm:"not null comment('用户名') index VARCHAR(100)"`
	Mobile            string    `json:"mobile" xorm:"not null comment('手机号') CHAR(11)"`
	GoodsId           int       `json:"goods_id" xorm:"not null comment('产品id') INT(11)"`
	PositionId        int       `json:"position_id" xorm:"not null default 0 comment('职位id') INT(11)"`
	ProductId         int       `json:"product_id" xorm:"not null comment('商品id') INT(11)"`
	GoodsPrice        string    `json:"goods_price" xorm:"not null comment('产品价格') DECIMAL(10,2)"`
	DiscountPrice     string    `json:"discount_price" xorm:"not null default 0.00 comment('优惠金额') DECIMAL(10,2)"`
	CourseId          int       `json:"course_id" xorm:"not null comment('课程id') INT(11)"`
	Price             string    `json:"price" xorm:"not null comment('支付金额') DECIMAL(10,2)"`
	Status            int       `json:"status" xorm:"not null default 0 comment('0 未支付 1已支付 3已取消 4已完成') index TINYINT(1)"`
	CreatedAt         time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"not null DATETIME"`
	PayTime           time.Time `json:"pay_time" xorm:"comment('支付时间') DATETIME"`
	OrderNum          string    `json:"order_num" xorm:"not null comment('订单编号') index VARCHAR(100)"`
	IsSync            int       `json:"is_sync" xorm:"not null default 0 comment('0 未同步 1同步神1成功') TINYINT(1)"`
	Title             string    `json:"title" xorm:"not null comment('课程标题') VARCHAR(150)"`
	PayRouteCode      string    `json:"pay_route_code" xorm:"VARCHAR(100)"`
	ChannelType       string    `json:"channel_type" xorm:"VARCHAR(100)"`
	PayChannelId      int       `json:"pay_channel_id" xorm:"INT(11)"`
	PayChannelName    string    `json:"pay_channel_name" xorm:"VARCHAR(100)"`
	EntryNodeId       int       `json:"entry_node_id" xorm:"INT(11)"`
	Sex               int       `json:"sex" xorm:"not null comment('0 女 1男') TINYINT(1)"`
	Uname             string    `json:"uname" xorm:"not null comment('用户填写用户名') VARCHAR(100)"`
	Cardid            string    `json:"cardid" xorm:"not null VARCHAR(100)"`
	Email             string    `json:"email" xorm:"not null VARCHAR(100)"`
	Zstype            int       `json:"zstype" xorm:"not null comment('0-不住宿 1-住宿') TINYINT(1)"`
	OrderId           int       `json:"order_id" xorm:"INT(11)"`
	Aid               int       `json:"aid" xorm:"not null comment('获取小能id') INT(11)"`
	DeviceType        string    `json:"device_type" xorm:"comment('设备类型') VARCHAR(20)"`
	DepartmentId      string    `json:"department_id" xorm:"comment('分部id') VARCHAR(20)"`
	DepartmentEhrCode string    `json:"department_ehr_code" xorm:"comment('分部ehr编码') VARCHAR(50)"`
	DepartmentName    string    `json:"department_name" xorm:"comment('分部名称') VARCHAR(50)"`
	SchoolId          string    `json:"school_id" xorm:"comment('分校id') VARCHAR(20)"`
	SchoolEhrCode     string    `json:"school_ehr_code" xorm:"comment('分校ehr编码') VARCHAR(50)"`
	SchoolName        string    `json:"school_name" xorm:"comment('分校名称') VARCHAR(50)"`
	CancelType        int       `json:"cancel_type" xorm:"comment('取消类型,1订单超时未支付取消，2用户主动取消') TINYINT(1)"`
	CancelReason      string    `json:"cancel_reason" xorm:"comment('取消原因') VARCHAR(50)"`
	OaoId             int       `json:"oao_id" xorm:"comment('教育oao课程id') INT(11)"`
	CourseType        int       `json:"course_type" xorm:"comment('课程类型：0面授课 1oao课') TINYINT(1)"`
	CancelAt          time.Time `json:"cancel_at" xorm:"comment('订单取消时间') DATETIME"`
	Province          string    `json:"province" xorm:"comment('报名省') VARCHAR(10)"`
	City              string    `json:"city" xorm:"comment('报名市') VARCHAR(10)"`
	Area              string    `json:"area" xorm:"comment('报名县区') VARCHAR(20)"`
	Address           string    `json:"address" xorm:"comment('报名详细地址') VARCHAR(50)"`
}
