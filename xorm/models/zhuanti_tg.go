package models

type ZhuantiTg struct {
	Tid            int    `json:"tid" xorm:"not null pk autoincr comment('表单id') INT(10)"`
	Flag           string `json:"flag" xorm:"comment('标记one.two.excel') VARCHAR(20)"`
	Filename       string `json:"fileName" xorm:"comment('专题名称') VARCHAR(50)"`
	Bigtitleone    string `json:"bigTitleOne" xorm:"comment('图文一大标题') VARCHAR(150)"`
	Netclasstitle  string `json:"netclassTitle" xorm:"comment('课程标题') VARCHAR(150)"`
	Netclassid     int    `json:"netclassId" xorm:"not null comment('课程id') INT(11)"`
	Litpic         string `json:"litpic" xorm:"comment('课程广告图') VARCHAR(255)"`
	Ishidebutton   int    `json:"isHideButton" xorm:"not null comment('按钮是否隐藏') INT(11)"`
	Buttontext     string `json:"buttonText" xorm:"comment('按钮文本') VARCHAR(20)"`
	Introduce      string `json:"introduce" xorm:"comment('课程介绍') VARCHAR(255)"`
	Bigtitletwo    string `json:"bigTitleTwo" xorm:"comment('图文二大标题') VARCHAR(255)"`
	Netclassid1    int    `json:"netclassId1" xorm:"not null comment('课程一id') INT(11)"`
	Netclasstitle1 string `json:"netclassTitle1" xorm:"comment('课程一名称') VARCHAR(255)"`
	Litpic1        string `json:"litpic1" xorm:"comment('课程一广告图') VARCHAR(255)"`
	Introduce1     string `json:"introduce1" xorm:"comment('课程一介绍') VARCHAR(255)"`
	Netclasstitle2 string `json:"netclassTitle2" xorm:"comment('课程二名称') VARCHAR(255)"`
	Netclassid2    int    `json:"netclassId2" xorm:"not null comment('课程二id') INT(11)"`
	Litpic2        string `json:"litpic2" xorm:"comment('课程二广告图') VARCHAR(255)"`
	Introduce2     string `json:"introduce2" xorm:"comment('课程二介绍') VARCHAR(255)"`
	Exceltitle     string `json:"excelTitle" xorm:"comment('表格名称') VARCHAR(255)"`
	Excel          string `json:"excel" xorm:"comment('表格内容') TEXT"`
}
