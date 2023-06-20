package models

type Synctbuser struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Username  string `json:"username" xorm:"not null VARCHAR(255)"`
	Type      int    `json:"type" xorm:"not null comment('1：网校余额积分同步用户 2：网校代金券同步用户 3：在线金币银币同步用户') INT(255)"`
	Usermoney string `json:"usermoney" xorm:"DECIMAL(10,2)"`
	Userpoint int    `json:"userpoint" xorm:"INT(11)"`
}
