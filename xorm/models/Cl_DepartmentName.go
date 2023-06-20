package models

type ClDepartmentname struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Departmentname string `json:"departmentName" xorm:"VARCHAR(50)"`
}
