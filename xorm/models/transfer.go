package models

import (
	"time"
)

type Transfer struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OldUserId    int64     `json:"old_user_id" xorm:"not null comment('原账号id') BIGINT(20)"`
	NewUserId    int64     `json:"new_user_id" xorm:"not null comment('新账号id') BIGINT(20)"`
	TransferName string    `json:"transfer_name" xorm:"not null comment('原账号姓名') VARCHAR(20)"`
	HandoverName string    `json:"handover_name" xorm:"not null comment('新账号姓名') VARCHAR(20)"`
	TransferMemo string    `json:"transfer_memo" xorm:"not null comment('移交备注') TEXT"`
	Operator     string    `json:"operator" xorm:"not null comment('操作人') VARCHAR(20)"`
	Status       int       `json:"status" xorm:"not null comment('移交状态是否已转	') TINYINT(1)"`
	TransferTime time.Time `json:"transfer_time" xorm:"not null comment('移交时间') DATETIME"`
	OldParentid  int       `json:"old_parentId" xorm:"not null comment('原账号的父级') INT(10)"`
	OldSubset    string    `json:"old_subset" xorm:"not null comment('原账号的子集') TEXT"`
	NewParentid  int       `json:"new_parentId" xorm:"not null comment('新账号的父级') INT(11)"`
	NewSubset    string    `json:"new_subset" xorm:"not null comment('新账号的子集') TEXT"`
}
