package models

import (
	"time"
)

type ShuatiOptions struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OptionName  string    `json:"option_name" xorm:"not null comment('配置项名称') index VARCHAR(64)"`
	OptionValue string    `json:"option_value" xorm:"not null comment('配置项值') TEXT"`
	CreatedAt   time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"not null DATETIME"`
}
