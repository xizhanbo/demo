package models

import (
	"time"
)

type SsGoods struct {
	Id                int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Goodsid           int       `json:"goodsId" xorm:"comment('商品id') index INT(11)"`
	Goodsname         string    `json:"goodsName" xorm:"default '' comment('商品名称') VARCHAR(255)"`
	Goodsno           string    `json:"goodsNo" xorm:"default '' comment('商品编号') VARCHAR(100)"`
	Goodsprice        int       `json:"goodsPrice" xorm:"comment('价格') INT(11)"`
	Courseid          int       `json:"courseId" xorm:"comment('所属课程id') INT(11)"`
	Branchschoolid    int       `json:"branchSchoolId" xorm:"comment('分校id') INT(11)"`
	Branchschoolname  string    `json:"branchSchoolName" xorm:"default '' comment('分校名称') VARCHAR(255)"`
	Subdepartmentid   int       `json:"subDepartmentId" xorm:"comment('分部id') index INT(11)"`
	Subdepartmentname string    `json:"subDepartmentName" xorm:"default '' comment('分部名称') VARCHAR(255)"`
	Starttime         string    `json:"startTime" xorm:"default '' comment('开课日期') VARCHAR(13)"`
	Endtime           string    `json:"endTime" xorm:"default '' comment('结束日期') VARCHAR(13)"`
	Createtime        string    `json:"createTime" xorm:"comment('创建日期') VARCHAR(13)"`
	Type              int       `json:"type" xorm:"default 1 comment('1:神一2：合伙人') TINYINT(1)"`
	Updatetime        time.Time `json:"updateTime" xorm:"comment('修改日期') DATETIME"`
	Buystatus         int       `json:"buyStatus" xorm:"default 1 comment('售卖状态:0禁止1允许') TINYINT(1)"`
}
