package models

import (
	"time"
)

type Zhuanti struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Type           string    `json:"type" xorm:"not null comment('专题风格(4种)') VARCHAR(50)"`
	Filename       string    `json:"fileName" xorm:"not null comment('专题文件名称') VARCHAR(50)"`
	Title          string    `json:"title" xorm:"not null comment('专题名') VARCHAR(150)"`
	Keywords       string    `json:"keywords" xorm:"not null comment('网页的核心搜索关键词') VARCHAR(300)"`
	Description    string    `json:"description" xorm:"not null comment('网页的描述信息') VARCHAR(500)"`
	Headtype       int       `json:"headType" xorm:"comment('头图内容修改类别(1为文字,2为图片)') INT(11)"`
	Smalltitle     string    `json:"smallTitle" xorm:"not null default '' comment('小标题') VARCHAR(150)"`
	Teacher        string    `json:"teacher" xorm:"not null comment('教师名') VARCHAR(150)"`
	Maintitle      string    `json:"mainTitle" xorm:"not null comment('主标题') VARCHAR(150)"`
	Subtitle       string    `json:"subtitle" xorm:"not null comment('副标题') VARCHAR(150)"`
	Teacherphoto   string    `json:"teacherPhoto" xorm:"not null comment('教师头像') VARCHAR(300)"`
	Mainpicture    string    `json:"mainPicture" xorm:"not null comment('头图图片') VARCHAR(300)"`
	Netclasstype   string    `json:"netclassType" xorm:"not null comment('课程推广形式(one:一个课程,two:两个课程,excel:表格)') VARCHAR(20)"`
	Bigtitleone    string    `json:"bigTitleOne" xorm:"not null comment('课程推荐大标题(一个课程)') VARCHAR(150)"`
	Netclasstitle  string    `json:"netclassTitle" xorm:"not null comment('课程标题(一个课程)') VARCHAR(150)"`
	Netclassid     int       `json:"netclassId" xorm:"comment('课程ID(一个课程)') INT(11)"`
	Litpic         string    `json:"litpic" xorm:"not null comment('课程广告图(一个课程)') VARCHAR(300)"`
	Buttontext     string    `json:"buttonText" xorm:"not null comment('按钮文字(一个课程)') VARCHAR(20)"`
	Ishidebutton   int       `json:"isHideButton" xorm:"comment('是否隐藏按钮 0:显示,1:隐藏(一个课程)') INT(11)"`
	Introduce      string    `json:"introduce" xorm:"not null comment('课程介绍(一个课程)') VARCHAR(300)"`
	Bigtitletwo    string    `json:"bigTitleTwo" xorm:"not null comment('课程推荐大标题(两个课程)') VARCHAR(300)"`
	Netclasstitle1 string    `json:"netclassTitle1" xorm:"not null comment('课程1标题(两个课程)') VARCHAR(150)"`
	Netclassid1    int       `json:"netclassId1" xorm:"comment('课程1ID(两个课程)') INT(11)"`
	Litpic1        string    `json:"litpic1" xorm:"not null comment('课程1广告图(两个课程)') VARCHAR(300)"`
	Introduce1     string    `json:"introduce1" xorm:"not null comment('课程1介绍(两个课程)') VARCHAR(300)"`
	Netclasstitle2 string    `json:"netclassTitle2" xorm:"not null comment('课程2标题(两个课程)') VARCHAR(150)"`
	Netclassid2    int       `json:"netclassId2" xorm:"comment('课程2ID(两个课程)') INT(11)"`
	Litpic2        string    `json:"litpic2" xorm:"not null comment('课程2广告图(两个课程)') VARCHAR(300)"`
	Introduce2     string    `json:"introduce2" xorm:"not null comment('课程2介绍(两个课程)') VARCHAR(300)"`
	Exceltitle     string    `json:"excelTitle" xorm:"not null comment('表格标题') VARCHAR(150)"`
	Excel          string    `json:"excel" xorm:"comment('表格形式显示推荐课程') TEXT"`
	Ishidezt       int       `json:"isHideZt" xorm:"comment('是否隐藏推荐专题 0:显示,1:隐藏') INT(11)"`
	Bigtitlezt     string    `json:"bigTitleZt" xorm:"not null comment('专题推荐大标题') VARCHAR(150)"`
	Titlezt1       string    `json:"titleZt1" xorm:"not null comment('专题活动1标题') VARCHAR(150)"`
	Urlzt1         string    `json:"urlZt1" xorm:"not null comment('专题活动1地址') VARCHAR(150)"`
	Litpiczt1      string    `json:"litpicZt1" xorm:"not null comment('专题活动1图片地址') VARCHAR(150)"`
	Titlezt2       string    `json:"titleZt2" xorm:"not null comment('专题活动2标题') VARCHAR(150)"`
	Urlzt2         string    `json:"urlZt2" xorm:"not null comment('专题活动2地址') VARCHAR(150)"`
	Litpiczt2      string    `json:"litpicZt2" xorm:"not null comment('专题活动2图片地址') VARCHAR(150)"`
	Titlezt3       string    `json:"titleZt3" xorm:"not null comment('专题活动3标题') VARCHAR(150)"`
	Urlzt3         string    `json:"urlZt3" xorm:"not null comment('专题活动3地址') VARCHAR(150)"`
	Litpiczt3      string    `json:"litpicZt3" xorm:"not null comment('专题活动3图片地址') VARCHAR(150)"`
	Plotter        string    `json:"plotter" xorm:"not null comment('创建人或策划人姓名') VARCHAR(50)"`
	Createdate     time.Time `json:"createDate" xorm:"comment('创建时间') DATETIME"`
	Lastdate       time.Time `json:"lastDate" xorm:"comment('最后一次修改时间') DATETIME"`
	Status         int       `json:"status" xorm:"comment('状态 0:未生成,1:已生成,3:已删除,4:已下线') INT(11)"`
	Username       string    `json:"userName" xorm:"not null comment('添加人用户名') VARCHAR(50)"`
	Userid         int       `json:"userId" xorm:"comment('添加人用户ID') INT(10)"`
	Lastusername   string    `json:"lastUserName" xorm:"not null comment('最后一次操作人') VARCHAR(50)"`
	Lastuserid     int       `json:"lastUserId" xorm:"comment('最后一次操作人ID') INT(11)"`
	Code           string    `json:"code" xorm:"not null comment('专题统计代码') VARCHAR(100)"`
}
