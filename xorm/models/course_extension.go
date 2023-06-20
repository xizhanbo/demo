package models

import (
	"time"
)

type CourseExtension struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ClassId           int       `json:"class_id" xorm:"not null comment('课程id') index INT(11)"`
	BookIds           string    `json:"book_ids" xorm:"comment('图书id,连接') index VARCHAR(200)"`
	CreatorName       string    `json:"creator_name" xorm:"comment('创建人用户名') VARCHAR(100)"`
	CreatorId         int       `json:"creator_id" xorm:"comment('创建人id') INT(11)"`
	AssignmentsCharge int       `json:"assignments_charge" xorm:"not null default 0 comment('课后作业是否收费') TINYINT(4)"`
	MockExam          string    `json:"mock_exam" xorm:"default '' comment('关联模考') VARCHAR(200)"`
	CreatedAt         time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt         time.Time `json:"updated_at" xorm:"DATETIME"`
	IsLiveLessons     int       `json:"is_live_lessons" xorm:"not null default 0 comment('是否是直播带课：0否，1是') TINYINT(1)"`
	SubjectiveNum     int       `json:"subjective_num" xorm:"comment('主观题批改次数') INT(10)"`
	TreatyType        int       `json:"treaty_type" xorm:"comment('协议类型 1关联协议参数 2关联协议模板 【1新协议课，2旧协议课】') TINYINT(4)"`
	TemplateId        int       `json:"template_id" xorm:"comment('协议模板id') INT(11)"`
	QiyuesuoId        int64     `json:"qiyuesuo_id" xorm:"comment('契约锁模板id') BIGINT(20)"`
	ParamId           int       `json:"param_id" xorm:"comment('参数id') INT(11)"`
	Staffnum          string    `json:"staffNum" xorm:"comment('OA工号') index VARCHAR(200)"`
	OfflineStartTime  time.Time `json:"offline_start_time" xorm:"comment('线下课开课日期') DATETIME"`
	OfflineStopTime   time.Time `json:"offline_stop_time" xorm:"comment('线下课结课日期') DATETIME"`
	MemberPrice       string    `json:"member_price" xorm:"comment('会员价') DECIMAL(10,2)"`
	YunduoCourseid    string    `json:"yunduo_courseId" xorm:"default '0' comment('crm主键id') VARCHAR(50)"`
	AchievementPrice  string    `json:"achievement_price" xorm:"comment('业绩结算金额') DECIMAL(10,2)"`
	OrderAsyncTime    time.Time `json:"order_async_time" xorm:"comment('订单同步时间') DATETIME"`
}
