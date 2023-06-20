package models

type UserGoodLessonsDataCount struct {
	Id        int    `json:"id" xorm:"not null pk autoincr comment('主键ID') INT(11)"`
	CateId    int    `json:"cate_id" xorm:"not null default 0 comment('分类ID') INT(11)"`
	UserName  string `json:"user_name" xorm:"not null default '' comment('用户名') index VARCHAR(50)"`
	Nums      int    `json:"nums" xorm:"not null default 0 comment('用户观看名师好课个数') INT(11)"`
	Frequency int    `json:"frequency" xorm:"not null default 0 comment('用户观看名师好课次数') INT(11)"`
	Status    int    `json:"status" xorm:"not null default 0 comment('状态 0有效 1无效') TINYINT(1)"`
}
