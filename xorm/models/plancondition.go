package models

type Plancondition struct {
	Id          int    `json:"id" xorm:"not null pk autoincr unique INT(10)"`
	Userid      int    `json:"userid" xorm:"INT(10)"`
	Nums        int    `json:"nums" xorm:"TINYINT(4)"`
	Datanum     int    `json:"dataNum" xorm:"TINYINT(4)"`
	Course      string `json:"course" xorm:"VARCHAR(255)"`
	Category    int    `json:"category" xorm:"TINYINT(4)"`
	Endexamdate int    `json:"endExamDate" xorm:"comment('考试结束时间') INT(10)"`
	Createtime  int    `json:"createTime" xorm:"INT(10)"`
}
