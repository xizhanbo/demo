package models

type HignendNum struct {
	Id              int     `json:"id" xorm:"not null pk autoincr INT(10)"`
	AreaName        string  `json:"area_name" xorm:"comment('省份名称') index VARCHAR(20)"`
	CompanyName     string  `json:"company_name" xorm:"comment('单位名称') index VARCHAR(50)"`
	CompanyNum      string  `json:"company_num" xorm:"comment('单位代码') index VARCHAR(30)"`
	PostName        string  `json:"post_name" xorm:"comment('岗位名称') VARCHAR(50)"`
	PostNum         string  `json:"post_num" xorm:"comment('岗位代码') VARCHAR(30)"`
	AdmissionNum    int     `json:"admission_num" xorm:"comment('招聘人数') TINYINT(3)"`
	Name            string  `json:"name" xorm:"comment('姓名') index VARCHAR(50)"`
	Admissionticket string  `json:"admissionticket" xorm:"comment('准考证') index VARCHAR(50)"`
	LinetestScore   float32 `json:"linetest_score" xorm:"comment('行测成绩') FLOAT(4,2)"`
	ShentestScore   float32 `json:"shentest_score" xorm:"comment('申论成绩') FLOAT(4,2)"`
	MinScore        float64 `json:"min_score" xorm:"comment('省份最低分数') DOUBLE"`
	AllSum          string  `json:"all_sum" xorm:"comment('总成绩/平均分') index VARCHAR(10)"`
	AllSort         int     `json:"all_sort" xorm:"comment('排名') INT(11)"`
}
