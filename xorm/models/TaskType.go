package models

type Tasktype struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('任务类型自增id') INT(11)"`
	Typename string `json:"typeName" xorm:"not null comment('类型名称') VARCHAR(255)"`
}
