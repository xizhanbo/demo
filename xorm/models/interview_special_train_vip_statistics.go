package models





type InterviewSpecialTrainVipStatistics struct {

	Id	int `json:"id" xorm:"not null pk autoincr INT(11)"`
	DateTime	string `json:"date_time" xorm:"not null default '' CHAR(10)"`
	2018PeopleCount	int `json:"2018_people_count" xorm:"not null default 0 INT(11)"`
	2018Achievement	string `json:"2018_achievement" xorm:"not null default 0.00 comment('2018业绩') DECIMAL(10,2)"`
	2019PeopleCount	int `json:"2019_people_count" xorm:"not null default 0 INT(11)"`
	2019Achievement	string `json:"2019_achievement" xorm:"not null default 0.00 comment('2019业绩') DECIMAL(10,2)"`
	2020PeopleCount	int `json:"2020_people_count" xorm:"not null default 0 INT(11)"`
	2020Achievement	string `json:"2020_achievement" xorm:"not null default 0.00 comment('2020业绩') DECIMAL(10,2)"`
	Status	int `json:"status" xorm:"not null default 0 comment('此条记录是否已经更新过') TINYINT(1)"`

}


