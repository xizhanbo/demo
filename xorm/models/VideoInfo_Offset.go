package models

type VideoinfoOffset struct {
	Offset int `json:"offset" xorm:"comment('标志位') INT(11)"`
	Type   int `json:"type" xorm:"comment('1：课件视频 2：公考讲堂') INT(11)"`
}
