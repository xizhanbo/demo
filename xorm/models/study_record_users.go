package models

type StudyRecordUsers struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	UserId   int    `json:"user_id" xorm:"not null unique INT(11)"`
	Username string `json:"username" xorm:"not null index VARCHAR(50)"`
}
