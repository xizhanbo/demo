package models

import (
	"time"
)

type Outboundtask struct {
	Rid              int       `json:"rid" xorm:"not null pk autoincr comment('主键') INT(10)"`
	Taskname         string    `json:"taskName" xorm:"not null comment('任务名称') index VARCHAR(30)"`
	Createat         time.Time `json:"createAt" xorm:"not null comment('任务创建时间') index DATETIME"`
	Start            time.Time `json:"start" xorm:"comment('数据清理开始时间') DATETIME"`
	End              time.Time `json:"end" xorm:"comment('数据清理结束时间') DATETIME"`
	Taskstatus       int       `json:"taskStatus" xorm:"not null default 1 comment('任务状态；1:处理中；2:已完成；3:已取消；4:处理失败') TINYINT(1)"`
	Username         string    `json:"userName" xorm:"comment('创建任务的用户名') VARCHAR(30)"`
	Updateat         time.Time `json:"updateAt" xorm:"comment('状态更新时间') DATETIME"`
	Uploadfilename   string    `json:"uploadFileName" xorm:"comment('上传文件名称') VARCHAR(50)"`
	Downloadfilepath string    `json:"downloadFilePath" xorm:"comment('下载文件路径') VARCHAR(100)"`
	Errorcause       string    `json:"errorCause" xorm:"comment('任务失败原因') VARCHAR(20)"`
}
