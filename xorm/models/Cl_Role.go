package models

type ClRole struct {
	Id           int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Rolename     string `json:"RoleName" xorm:"VARCHAR(200)"`
	Rolepurview  string `json:"RolePurView" xorm:"VARCHAR(800)"`
	Optuserid    int    `json:"OptUserId" xorm:"INT(11)"`
	Optusername  string `json:"OptUserName" xorm:"VARCHAR(200)"`
	Departmentid int    `json:"DepartmentId" xorm:"INT(11)"`
}
