package models

import (
	"time"
)

type ProtocolClassRefund struct {
	Id               int       `json:"id" xorm:"not null pk autoincr comment('表ID') INT(11)"`
	OrderId          int       `json:"order_id" xorm:"not null default 0 comment('cl_order表ID') INT(11)"`
	ClassId          int       `json:"class_id" xorm:"not null default 0 comment('NetClasses表ID') INT(11)"`
	UserId           int       `json:"user_id" xorm:"not null default 0 comment('用户ID') INT(11)"`
	Score            string    `json:"score" xorm:"not null default '' comment('成绩分数') VARCHAR(10)"`
	Type             int       `json:"type" xorm:"not null default 1 comment('类型1笔试 2面试') TINYINT(1)"`
	State            int       `json:"state" xorm:"not null default 0 comment('退费状态：0申请成功，1处理中，2申请通过，3申请未通过/驳回，4退费中，5已退款') TINYINT(1)"`
	RecruitNumber    int       `json:"recruit_number" xorm:"not null default 0 comment('职位招考人数') INT(11)"`
	BankCardNumber   string    `json:"bank_card_number" xorm:"not null default '' comment('银行卡号') VARCHAR(100)"`
	BankUsername     string    `json:"bank_username" xorm:"not null default '' comment('银行卡户主名') VARCHAR(20)"`
	SubmitSource     int       `json:"submit_source" xorm:"not null default 1 comment('协议退费申请来源 ：１学员通过前端提交　２客服通过协议班退费流程提交') TINYINT(1)"`
	IdentityFrontPic string    `json:"identity_front_pic" xorm:"not null default '' comment('身份证正面图片URL') VARCHAR(255)"`
	IdentityBackPic  string    `json:"identity_back_pic" xorm:"not null default '' comment('身份证反面图片URL') VARCHAR(255)"`
	OlicyBonus       int       `json:"olicy_bonus" xorm:"not null default 0 comment('政策加分') INT(11)"`
	ExamCardPic      string    `json:"exam_card_pic" xorm:"not null default '' comment('准考证URL') VARCHAR(255)"`
	WrittenExamPic   string    `json:"written_exam_pic" xorm:"not null default '' comment('笔试成绩单URL') VARCHAR(255)"`
	InterviewPic     string    `json:"interview_pic" xorm:"not null default '' comment('面试成绩单URL') VARCHAR(255)"`
	EntryFacePic     string    `json:"entry_face_pic" xorm:"not null default '' comment('进面名单URL') VARCHAR(255)"`
	RefundFormUrl    string    `json:"refund_form_url" xorm:"not null default '' comment('客服提交的退费申请单URL') VARCHAR(255)"`
	PublicListPic    string    `json:"public_list_pic" xorm:"not null default '' comment('公示名单URL') VARCHAR(255)"`
	PhysicalListPic  string    `json:"physical_list_pic" xorm:"not null default '' comment('体检名单URL') VARCHAR(255)"`
	EmployListPic    string    `json:"employ_list_pic" xorm:"not null default '' comment('最终拟录用名单URL') VARCHAR(255)"`
	CreateAt         time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdateAt         time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}
