package models





type Zhaosheng struct {

	Id	int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Riqi	string `json:"riqi" xorm:"not null default '' comment('日期（一年内，每一天的日历）') CHAR(10)"`
	2015renshu	int `json:"2015renshu" xorm:"not null default 0 comment('2015年招生人数（当日已到账的月卡订单数+普通订单数）') INT(11)"`
	2015yeji	string `json:"2015yeji" xorm:"not null default 0.00 comment('2015年业绩（当日所有业绩）') DECIMAL(10,2)"`
	2016renshu	int `json:"2016renshu" xorm:"not null default 0 comment('2016年招生人数（当日已到账的月卡订单数+普通订单数）') INT(11)"`
	2016yeji	string `json:"2016yeji" xorm:"not null default 0.00 comment('2015年业绩（当日所有业绩）') DECIMAL(10,2)"`
	2016cashorders	int `json:"2016cashOrders" xorm:"not null default 0 comment('使用现金支付的订单(人)数') INT(11)"`
	2016monthcardorder	int `json:"2016monthCardOrder" xorm:"not null default 0 comment('月卡订单数') INT(11)"`
	2017renshu	int `json:"2017renshu" xorm:"not null default 0 INT(11)"`
	2017yeji	string `json:"2017yeji" xorm:"not null default 0 DECIMAL(10)"`
	2017cashorders	int `json:"2017cashOrders" xorm:"not null default 0 INT(11)"`
	2017monthcardorder	int `json:"2017monthCardOrder" xorm:"not null default 0 INT(11)"`
	2018renshu	int `json:"2018renshu" xorm:"not null default 0 INT(11)"`
	2018yeji	string `json:"2018yeji" xorm:"not null default 0 DECIMAL(10)"`
	2018cashorders	int `json:"2018cashOrders" xorm:"not null default 0 INT(11)"`
	2018monthcardorder	int `json:"2018monthCardOrder" xorm:"not null default 0 INT(11)"`
	2019renshu	int `json:"2019renshu" xorm:"not null default 0 INT(11)"`
	2019yeji	string `json:"2019yeji" xorm:"not null default 0 DECIMAL(10)"`
	2019cashorders	int `json:"2019cashOrders" xorm:"not null default 0 INT(11)"`
	2019monthcardorder	int `json:"2019monthCardOrder" xorm:"not null default 0 INT(11)"`
	2020renshu	int `json:"2020renshu" xorm:"not null default 0 INT(11)"`
	2020yeji	string `json:"2020yeji" xorm:"not null default 0 DECIMAL(10)"`
	2020cashorders	int `json:"2020cashOrders" xorm:"not null default 0 INT(11)"`
	2020monthcardorder	int `json:"2020monthCardOrder" xorm:"not null default 0 INT(11)"`
	2021renshu	int `json:"2021renshu" xorm:"not null default 0 INT(11)"`
	2021yeji	string `json:"2021yeji" xorm:"not null default 0 DECIMAL(10)"`
	2021cashorders	int `json:"2021cashOrders" xorm:"not null default 0 INT(11)"`
	2021monthcardorder	int `json:"2021monthCardOrder" xorm:"not null default 0 INT(11)"`
	2022renshu	int `json:"2022renshu" xorm:"not null default 0 INT(11)"`
	2022yeji	string `json:"2022yeji" xorm:"not null default 0 DECIMAL(10)"`
	2022cashorders	int `json:"2022cashOrders" xorm:"not null default 0 INT(11)"`
	2022monthcardorder	int `json:"2022monthCardOrder" xorm:"not null default 0 INT(11)"`
	2023renshu	int `json:"2023renshu" xorm:"not null default 0 INT(11)"`
	2023yeji	string `json:"2023yeji" xorm:"not null default 0 DECIMAL(10)"`
	2023cashorders	int `json:"2023cashOrders" xorm:"not null default 0 INT(11)"`
	2023monthcardorder	int `json:"2023monthCardOrder" xorm:"not null default 0 INT(11)"`
	2024renshu	int `json:"2024renshu" xorm:"not null default 0 INT(11)"`
	2024yeji	string `json:"2024yeji" xorm:"not null default 0 DECIMAL(10)"`
	2024cashorders	int `json:"2024cashOrders" xorm:"not null default 0 INT(11)"`
	2024monthcardorder	int `json:"2024monthCardOrder" xorm:"not null default 0 INT(11)"`
	2025renshu	int `json:"2025renshu" xorm:"not null default 0 INT(11)"`
	2025yeji	string `json:"2025yeji" xorm:"not null default 0 DECIMAL(10)"`
	2025cashorders	int `json:"2025cashOrders" xorm:"not null default 0 INT(11)"`
	2025monthcardorder	int `json:"2025monthCardOrder" xorm:"not null default 0 INT(11)"`
	2026renshu	int `json:"2026renshu" xorm:"not null default 0 INT(11)"`
	2026yeji	string `json:"2026yeji" xorm:"not null default 0 DECIMAL(10)"`
	2026cashorders	int `json:"2026cashOrders" xorm:"not null default 0 INT(11)"`
	2026monthcardorder	int `json:"2026monthCardOrder" xorm:"not null default 0 INT(11)"`
	2027renshu	int `json:"2027renshu" xorm:"not null default 0 INT(11)"`
	2027yeji	string `json:"2027yeji" xorm:"not null default 0 DECIMAL(10)"`
	2027cashorders	int `json:"2027cashOrders" xorm:"not null default 0 INT(11)"`
	2027monthcardorder	int `json:"2027monthCardOrder" xorm:"not null default 0 INT(11)"`
	Status	int `json:"status" xorm:"not null default 0 comment('此条记录是否已经更新过') TINYINT(1)"`

}


