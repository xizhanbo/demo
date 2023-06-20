package models

import (
	"time"
)

type InspectorStudentClass struct {
	Id               int       `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	AdminName        string    `json:"admin_name" xorm:"not null default '' comment('跟进人 【Cl_Admin表的UserName】') index CHAR(100)"`
	StudentName      string    `json:"student_name" xorm:"not null default '' comment('学员用户名【Cl_User表的UserName】') index(student_class_id) CHAR(100)"`
	ClassId          int       `json:"class_id" xorm:"not null default 0 comment('课程ID【NetClasses表的rid】') index index(student_class_id) INT(11)"`
	OrderId          int       `json:"order_id" xorm:"not null default 0 comment('课程的订单ID【Cl_Order表ID】') INT(11)"`
	Operator         string    `json:"operator" xorm:"not null default '' comment('操作人') VARCHAR(255)"`
	ClaimTime        time.Time `json:"claim_time" xorm:"not null comment('认领时间') index DATETIME"`
	PayTime          time.Time `json:"pay_time" xorm:"not null comment('订单付款时间【Cl_Order表的PayDate】') DATETIME"`
	WorkStatus       int       `json:"work_status" xorm:"not null default 0 comment('工作状态：0--，1在职，2全职') TINYINT(1)"`
	ExperienceStatus int       `json:"experience_status" xorm:"not null default 0 comment('备考经验：0--，1有，2无') TINYINT(1)"`
	WechatStatus     int       `json:"wechat_status" xorm:"not null comment('添加微信：0--，1是，2否') TINYINT(1)"`
	ClassGroup       string    `json:"class_group" xorm:"not null default '' comment('班级群') CHAR(50)"`
	Nickname         string    `json:"nickname" xorm:"not null default '' comment('微信昵称') CHAR(50)"`
	StudyPlanDate    time.Time `json:"study_plan_date" xorm:"comment('发学习计划时间【格式：年月日】') DATE"`
	ReportDate       time.Time `json:"report_date" xorm:"comment('给报告时间【格式：年月日】') DATE"`
	RecordText       string    `json:"record_text" xorm:"comment('学员主动咨询【不超过500个字】') TEXT"`
	UpdateAt         time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	CreateAt         time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}
