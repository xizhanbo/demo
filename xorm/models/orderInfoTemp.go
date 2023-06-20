package models

type Orderinfotemp struct {
	Id           int    `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	Classid      int    `json:"classId" xorm:"not null comment('课程ID') INT(11)"`
	Username     string `json:"username" xorm:"not null comment('用户名') index VARCHAR(60)"`
	Consignee    string `json:"consignee" xorm:"not null comment('收货人姓名') VARCHAR(60)"`
	Provinceid   string `json:"provinceid" xorm:"not null default '' comment('省份id') VARCHAR(6)"`
	Cityid       string `json:"cityid" xorm:"not null default '' comment('城市id') VARCHAR(6)"`
	Areaid       string `json:"areaid" xorm:"not null comment('地区id') VARCHAR(6)"`
	Provincename string `json:"provinceName" xorm:"comment('所属省份名称') VARCHAR(100)"`
	Cityname     string `json:"cityName" xorm:"comment('地区名称') VARCHAR(100)"`
	Areaname     string `json:"areaName" xorm:"comment('地区名称') VARCHAR(100)"`
	Address      string `json:"address" xorm:"not null comment('详细地址') VARCHAR(255)"`
	Phone        int64  `json:"phone" xorm:"not null comment('联系人手机号码') BIGINT(11)"`
	Isdefault    int    `json:"isDefault" xorm:"default 0 comment('是否是默认收货地址') TINYINT(1)"`
	Status       int    `json:"status" xorm:"default 1 comment('地址是否可用') TINYINT(1)"`
}
