package models

import (
	"time"
)

type StudentInfo struct {
	Id                   int64  `json:"id" xorm:"pk autoincr comment('主键 特殊意义') BIGINT(20)"`
	Phone                string `json:"phone" xorm:"not null comment(' 机号') VARCHAR(32)"`
	UName                string `json:"u_name" xorm:"comment('姓名') VARCHAR(32)"`
	Gender               string `json:"gender" xorm:"comment('性别: 男， ，未知') VARCHAR(32)"`
	Province             string `json:"province" xorm:"not null comment('省份') VARCHAR(32)"`
	City                 string `json:"city" xorm:"comment('地市') VARCHAR(32)"`
	District             string `json:"district" xorm:"comment('区县') VARCHAR(32)"`
	Expense              int    `json:"expense" xorm:"comment(' 额') INT(20)"`
	ExamType             string `json:"exam_type" xorm:"comment('考试类型') VARCHAR(32)"`
	Education            string `json:"education" xorm:"comment('学历') VARCHAR(32)"`
	GraduateInstitutions string `json:"graduate_institutions" xorm:"comment('毕业院
校') VARCHAR(32)"`
	ActionType string    `json:"action_type" xorm:"not null comment(' 户 为(购书、注册等)') VARCHAR(32)"`
	ActionTime time.Time `json:"action_time" xorm:"not null comment(' 为发 时间') DATETIME"`
	ImportWay  string    `json:"import_way" xorm:"not null comment('数据来源') VARCHAR(32)"`
	ImportTime time.Time `json:"import_time" xorm:"not null comment('本次收集时间') DATETIME"`
	Remark     string    `json:"remark" xorm:"comment('备注') VARCHAR(256)"`
}
