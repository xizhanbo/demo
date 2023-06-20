package models

type DzSchool struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	School string `json:"school" xorm:"not null default '' comment('学校名称') VARCHAR(100)"`
	Father int    `json:"father" xorm:"not null comment('父级') INT(11)"`
}
