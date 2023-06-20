package models

import (
	"time"
)

type Task struct {
	Tid            int       `json:"Tid" xorm:"not null pk autoincr comment('任务id（自增任务id唯一索引）') INT(11)"`
	Taskname       string    `json:"taskName" xorm:"comment('任务名称') VARCHAR(255)"`
	Tasktype       int       `json:"taskType" xorm:"comment('任务类型1网站前台2网站后台3接口4专题') INT(11)"`
	Requestpeople  string    `json:"requestPeople" xorm:"comment('需求方') VARCHAR(255)"`
	Dotaskpeople   string    `json:"doTaskPeople" xorm:"comment('做任务的人') VARCHAR(255)"`
	Dotaskusername string    `json:"doTaskUserName" xorm:"not null comment('接受任务用户名用，隔开') VARCHAR(50)"`
	Starttime      time.Time `json:"starttime" xorm:"comment('开始时间') DATETIME"`
	Expecttime     time.Time `json:"expecttime" xorm:"comment('预计完成时间') DATETIME"`
	Endtime        time.Time `json:"endtime" xorm:"comment('实际完成时间') DATETIME"`
	Remark         string    `json:"remark" xorm:"comment('备注') VARCHAR(255)"`
	Optuserid      int       `json:"optUserid" xorm:"comment('添加人ID') INT(11)"`
	Optuser        string    `json:"optUser" xorm:"comment('添加人用户名') VARCHAR(255)"`
	State          int       `json:"state" xorm:"default 1 comment('任务状态1已废弃2完成3进行中4未开始5待确认') INT(11)"`
	Devtype        int       `json:"devType" xorm:"comment('开发类型1前端2技术后端') INT(11)"`
	Weight         int       `json:"weight" xorm:"default 3 comment('1紧急且重要2紧急但不重要3不紧急但重要4不紧急不重要') INT(11)"`
}
