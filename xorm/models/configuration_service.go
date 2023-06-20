package models

import (
	"time"
)

type ConfigurationService struct {
	Id              int       `json:"id" xorm:"not null pk autoincr comment('主键id') INT(11)"`
	ServiceNum      int       `json:"service_num" xorm:"not null comment('服务号') INT(10)"`
	ChannelNo       string    `json:"channel_no" xorm:"not null default '0' comment('渠道号') unique VARCHAR(50)"`
	ChannelName     string    `json:"channel_name" xorm:"comment('渠道名') VARCHAR(50)"`
	QrCode          string    `json:"qr_code" xorm:"comment('二维码') VARCHAR(100)"`
	Status          int       `json:"status" xorm:"default 0 comment('状态：0可用，1删除') TINYINT(1)"`
	CreateName      string    `json:"create_name" xorm:"VARCHAR(50)"`
	CreateTime      time.Time `json:"create_time" xorm:"comment('创建时间') DATETIME"`
	CostomerMessage string    `json:"costomer_message" xorm:"comment('客服消息') TEXT"`
}
