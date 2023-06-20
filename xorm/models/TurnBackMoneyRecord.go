package models

import (
	"time"
)

type Turnbackmoneyrecord struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Orderitemid  int       `json:"OrderItemId" xorm:"not null INT(11)"`
	Isbank       int       `json:"IsBank" xorm:"not null default 0 comment('0 退款到账户，1，退款到银行') TINYINT(4)"`
	Bankname     string    `json:"BankName" xorm:"VARCHAR(200)"`
	Bankusername string    `json:"BankUserName" xorm:"VARCHAR(200)"`
	Province     string    `json:"province" xorm:"VARCHAR(100)"`
	City         string    `json:"city" xorm:"VARCHAR(100)"`
	Bankaccount  string    `json:"BankAccount" xorm:"VARCHAR(100)"`
	Oa           string    `json:"OA" xorm:"VARCHAR(100)"`
	Status       int       `json:"Status" xorm:"not null default 0 comment('0处理中 1完成 2取消') INT(11)"`
	Addtime      time.Time `json:"AddTime" xorm:"not null DATETIME"`
	Optuserid    int       `json:"OptUserID" xorm:"not null INT(11)"`
	Optusername  string    `json:"OptUserName" xorm:"not null VARCHAR(100)"`
	Statustime   time.Time `json:"StatusTime" xorm:"comment('状态时间') DATETIME"`
	Remark       string    `json:"Remark" xorm:"VARCHAR(300)"`
}
