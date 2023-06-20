package models





type Bishicount struct {

	Id	int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Riqi	string `json:"riqi" xorm:"not null default '' CHAR(10)"`
	2016renshu	int `json:"2016renshu" xorm:"not null default 0 INT(11)"`
	2016yeji	string `json:"2016yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2016cashorders	int `json:"2016cashOrders" xorm:"not null default 0 INT(11)"`
	2017renshu	int `json:"2017renshu" xorm:"not null default 0 INT(11)"`
	2017yeji	string `json:"2017yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2017cashorders	int `json:"2017cashOrders" xorm:"not null default 0 INT(11)"`
	2018renshu	int `json:"2018renshu" xorm:"not null default 0 INT(11)"`
	2018yeji	string `json:"2018yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2018cashorders	int `json:"2018cashOrders" xorm:"not null default 0 INT(11)"`
	2019renshu	int `json:"2019renshu" xorm:"not null default 0 INT(11)"`
	2019yeji	string `json:"2019yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2019cashorders	int `json:"2019cashOrders" xorm:"not null default 0 INT(11)"`
	2020renshu	int `json:"2020renshu" xorm:"not null default 0 INT(11)"`
	2020yeji	string `json:"2020yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2020cashorders	int `json:"2020cashOrders" xorm:"not null default 0 INT(11)"`
	2021renshu	int `json:"2021renshu" xorm:"not null default 0 INT(11)"`
	2021yeji	string `json:"2021yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2021cashorders	int `json:"2021cashOrders" xorm:"not null default 0 INT(11)"`
	2022renshu	int `json:"2022renshu" xorm:"not null default 0 INT(11)"`
	2022yeji	string `json:"2022yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2022cashorders	int `json:"2022cashOrders" xorm:"not null default 0 INT(11)"`
	2023renshu	int `json:"2023renshu" xorm:"not null default 0 INT(11)"`
	2023yeji	string `json:"2023yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2023cashorders	int `json:"2023cashOrders" xorm:"not null default 0 INT(11)"`
	2024renshu	int `json:"2024renshu" xorm:"not null default 0 INT(11)"`
	2024yeji	string `json:"2024yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2024cashorders	int `json:"2024cashOrders" xorm:"not null default 0 INT(11)"`
	2025renshu	int `json:"2025renshu" xorm:"not null default 0 INT(11)"`
	2025yeji	string `json:"2025yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2025cashorders	int `json:"2025cashOrders" xorm:"not null default 0 INT(11)"`
	2026renshu	int `json:"2026renshu" xorm:"not null default 0 INT(11)"`
	2026yeji	string `json:"2026yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2026cashorders	int `json:"2026cashOrders" xorm:"not null default 0 INT(11)"`
	2027renshu	int `json:"2027renshu" xorm:"not null default 0 INT(11)"`
	2027yeji	string `json:"2027yeji" xorm:"not null default 0.00 DECIMAL(10,2)"`
	2027cashorders	int `json:"2027cashOrders" xorm:"not null default 0 INT(11)"`
	Status	int `json:"status" xorm:"not null default 0 comment('此条记录是否已经更新过') TINYINT(1)"`

}


