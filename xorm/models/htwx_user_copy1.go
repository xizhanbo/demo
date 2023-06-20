package models

type HtwxUserCopy1 struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UcId     int    `json:"uc_id" xorm:"INT(11)"`
	Phone    string `json:"phone" xorm:"index VARCHAR(255)"`
	Username string `json:"username" xorm:"index VARCHAR(255)"`
	Userid   int    `json:"userid" xorm:"unique INT(11)"`
	Status   int    `json:"status" xorm:"default 0 INT(11)"`
}
