package models

type ClDepartment struct {
	Id             int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Departmentname string `json:"DepartmentName" xorm:"VARCHAR(50)"`
	Status         int    `json:"status" xorm:"default 0 INT(11)"`
}
