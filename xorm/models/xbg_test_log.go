package models

type XbgTestLog struct {
	Id   int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Text string `json:"text" xorm:"TEXT"`
}
