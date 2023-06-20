package models

type Plan struct {
	Id            int    `json:"id" xorm:"not null pk autoincr INT(8)"`
	Uid           int    `json:"uid" xorm:"not null INT(8)"`
	Lessionid     int    `json:"lessionid" xorm:"default 0 INT(8)"`
	Classid       int    `json:"classid" xorm:"default 0 INT(8)"`
	Topid         int    `json:"topid" xorm:"default 0 INT(8)"`
	Lestitle      string `json:"lestitle" xorm:"comment('课件名字') VARCHAR(255)"`
	Netclasstitle string `json:"netclasstitle" xorm:"comment('班次名字') VARCHAR(255)"`
	Plandate      string `json:"plandate" xorm:"comment('学习时间') VARCHAR(255)"`
	Type          int    `json:"type" xorm:"default 1 comment('1 课程 2 试卷 3 直播') TINYINT(1)"`
	Monthcourse   int    `json:"monthcourse" xorm:"default 0 comment('月卡表示') TINYINT(1)"`
	Videotype     int    `json:"videotype" xorm:"default 0 comment('直播 录播表示') TINYINT(1)"`
	Netclasstype  int    `json:"netclasstype" xorm:"default 0 comment('铺子表示') TINYINT(1)"`
	Teacherid     int    `json:"teacherid" xorm:"default 0 INT(8)"`
	Teachername   string `json:"teachername" xorm:"comment('教师姓名') VARCHAR(255)"`
	Status        int    `json:"status" xorm:"default 0 comment('学习 1 未学习 0') TINYINT(1)"`
	Plantype      int    `json:"planType" xorm:"default 1 comment('1 日历使用 2 代表普通日历使用') TINYINT(1)"`
	Timelength    int    `json:"timelength" xorm:"default 0 INT(8)"`
	Textnum       int    `json:"textNum" xorm:"default 0 comment('题目数量') SMALLINT(4)"`
	Shorttitle    string `json:"shorttitle" xorm:"comment('短标题') VARCHAR(255)"`
	Createtime    int    `json:"createtime" xorm:"INT(10)"`
	Category      int    `json:"category" xorm:"comment('考试类型') TINYINT(2)"`
	Begintime     string `json:"beginTime" xorm:"default '0' comment('直播开始时间') VARCHAR(50)"`
	Endtime       string `json:"endTime" xorm:"default '0' comment('直播结束时间') VARCHAR(50)"`
	Paperid       int    `json:"paperId" xorm:"comment('试卷id') INT(10)"`
	Ordernum      int    `json:"ordernum" xorm:"comment('排序') SMALLINT(4)"`
}
