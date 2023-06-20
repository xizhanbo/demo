package models

type Diqu struct {
	Id      int    `json:"id" xorm:"not null pk comment('省市县代表的ID值') index INT(11)"`
	Name    string `json:"name" xorm:"comment('省市县名称') VARCHAR(50)"`
	Father  int    `json:"father" xorm:"comment('父级ID值') index INT(11)"`
	Type    string `json:"type" xorm:"comment('类型，1省，2市，3县') VARCHAR(20)"`
	Acronym string `json:"acronym" xorm:"comment('第一个字的首字母') VARCHAR(5)"`
}
