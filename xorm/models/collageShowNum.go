package models

type Collageshownum struct {
	Id         int64 `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Activityid int   `json:"activityId" xorm:"not null comment('拼团活动id') unique INT(11)"`
	Shownum    int   `json:"showNum" xorm:"comment('已拼销量') INT(11)"`
}
