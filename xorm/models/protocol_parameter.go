package models

import (
	"time"
)

type ProtocolParameter struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ParamName         string    `json:"param_name" xorm:"not null comment('协议名称') index VARCHAR(200)"`
	TemplateType      int       `json:"template_type" xorm:"not null comment('协议模板 protocol_template表ID') INT(11)"`
	CategoryId        int       `json:"category_id" xorm:"not null comment('考试类型') INT(11)"`
	ExamName          string    `json:"exam_name" xorm:"not null comment('考试名称') index VARCHAR(200)"`
	ClassName         string    `json:"class_name" xorm:"not null default '' comment('班级名称') VARCHAR(200)"`
	ProtocolType      int       `json:"protocol_type" xorm:"not null comment('协议类别 1笔试 2面试') INT(11)"`
	ClassCost         string    `json:"class_cost" xorm:"not null comment('课程成本（元）') index DECIMAL(10,2)"`
	PayType           int       `json:"pay_type" xorm:"not null default 1 comment('付费模式 1全款2预收') TINYINT(4)"`
	CreateName        string    `json:"create_name" xorm:"not null default '' comment('创建人') VARCHAR(50)"`
	CreateAt          time.Time `json:"create_at" xorm:"not null comment('创建时间') DATETIME"`
	UpdateName        string    `json:"update_name" xorm:"not null default '' comment('修改人') VARCHAR(50)"`
	UpdateAt          time.Time `json:"update_at" xorm:"not null comment('修改时间') DATETIME"`
	AdvancePriceSmall string    `json:"advance_price_small" xorm:"comment('预支付培训费（小写）') VARCHAR(100)"`
	AdvancePriceBig   string    `json:"advance_price_big" xorm:"comment('预支付培训费（大写）') VARCHAR(100)"`
	BackPriceSmall    string    `json:"back_price_small" xorm:"comment('补交培训费原价（小写）') VARCHAR(100)"`
	BackPriceBig      string    `json:"back_price_big" xorm:"comment('补交培训费原价（大写）') VARCHAR(100)"`
	BackDiscountSmall string    `json:"back_discount_small" xorm:"comment('补交培训费优惠价（小写）') VARCHAR(100)"`
	BackDiscountBig   string    `json:"back_discount_big" xorm:"comment('补交培训费优惠价（大写）') VARCHAR(100)"`
	HourPrice         string    `json:"hour_price" xorm:"comment('线上课时费') VARCHAR(100)"`
	UnitPrice         string    `json:"unit_price" xorm:"comment('每日课单价金额') VARCHAR(100)"`
	StudyNeeds        string    `json:"study_needs" xorm:"comment('学习要求') TEXT"`
	Remark1           string    `json:"remark1" xorm:"comment('自定义1') VARCHAR(100)"`
	Remark2           string    `json:"remark2" xorm:"comment('自定义2') VARCHAR(100)"`
	Remark3           string    `json:"remark3" xorm:"comment('自定义3') VARCHAR(100)"`
	Remark4           string    `json:"remark4" xorm:"comment('自定义4') VARCHAR(100)"`
	Remark5           string    `json:"remark5" xorm:"comment('自定义5') VARCHAR(100)"`
	Remark6           string    `json:"remark6" xorm:"comment('自定义6') VARCHAR(100)"`
	RefundType        int       `json:"refund_type" xorm:"not null default 0 comment('退款方式 1全额退款 2扣除实收') TINYINT(3)"`
}
