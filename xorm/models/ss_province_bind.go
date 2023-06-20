package models

import (
	"time"
)

type SsProvinceBind struct {
	Id               int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ProvinceId       int       `json:"province_id" xorm:"comment('省份id') index(province_idx) INT(11)"`
	ProvinceName     string    `json:"province_name" xorm:"comment('省份名称') index(province_idx) VARCHAR(255)"`
	BranchSchoolId   int       `json:"branch_school_id" xorm:"comment('分校id') index(branch_idx) INT(11)"`
	BranchSchoolName string    `json:"branch_school_name" xorm:"comment('分校名称') index(branch_idx) VARCHAR(255)"`
	Operator         string    `json:"operator" xorm:"comment('添加人') VARCHAR(255)"`
	Addtime          time.Time `json:"addtime" xorm:"comment('添加时间') DATETIME"`
}
