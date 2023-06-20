package models

type Api struct {
	Id        int    `json:"id" xorm:"not null pk autoincr comment('接口编号') INT(11)"`
	Aid       int    `json:"aid" xorm:"default 0 comment('接口分类id') INT(11)"`
	Num       string `json:"num" xorm:"comment('接口编号') VARCHAR(100)"`
	Url       string `json:"url" xorm:"comment('请求地址') VARCHAR(240)"`
	Name      string `json:"name" xorm:"comment('接口名') VARCHAR(100)"`
	Des       string `json:"des" xorm:"comment('接口描述') VARCHAR(300)"`
	Parameter string `json:"parameter" xorm:"comment('请求参数{所有的主求参数,以json格式在此存放}') TEXT"`
	Memo      string `json:"memo" xorm:"comment('备注') TEXT"`
	Re        string `json:"re" xorm:"comment('返回值') TEXT"`
	Lasttime  int    `json:"lasttime" xorm:"comment('提后操作时间') INT(11)"`
	Lastuid   int    `json:"lastuid" xorm:"comment('最后修改uid') INT(11)"`
	Creatuid  int    `json:"creatuid" xorm:"comment('创建者的UID') INT(11)"`
	Isdel     int    `json:"isdel" xorm:"default 0 comment('{0:正常,1:下线,2:删除}') TINYINT(4)"`
	Type      string `json:"type" xorm:"comment('请求方式') CHAR(11)"`
	Ord       int    `json:"ord" xorm:"default 0 comment('排序(值越大,越靠前)') INT(11)"`
}
