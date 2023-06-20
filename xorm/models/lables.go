package models

import (
	"time"
)

type Lables struct {
	Id        int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name      string    `json:"name" xorm:"not null default '' VARCHAR(30)"`
	Type      int       `json:"type" xorm:"not null default 0 TINYINT(1)"`
	Attribute int       `json:"attribute" xorm:"not null default 0 TINYINT(1)"`
	Sort      int       `json:"sort" xorm:"not null default 0 INT(11)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	DeletedAt time.Time `json:"deleted_at" xorm:"DATETIME"`
}
