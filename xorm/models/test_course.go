package models

type TestCourse struct {
	Id           int `json:"id" xorm:"not null pk autoincr INT(10)"`
	Courseid     int `json:"courseid" xorm:"not null default 0 INT(10)"`
	Netclassid   int `json:"netclassid" xorm:"not null default 0 INT(10)"`
	Classpackage int `json:"classPackage" xorm:"not null default 0 INT(10)"`
	Paperid      int `json:"paperid" xorm:"not null INT(10)"`
	Type         int `json:"type" xorm:"not null TINYINT(1)"`
	Position     int `json:"position" xorm:"not null comment('1:课前,2:课后') TINYINT(3)"`
	Status       int `json:"status" xorm:"not null default 1 TINYINT(1)"`
}
