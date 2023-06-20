package models

import (
	"time"
)

type MsgMessage struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Type     int    `json:"type" xorm:"not null comment('消息类别  1:活动 2:课程 (新增3:消息)') TINYINT(1)"`
	Category string `json:"category" xorm:"not null default '' comment('考试类别
 999=>'不限',1=>'公务员考试',2=>'公务员面试',3=>'社会工作师考试',4=>'招警考试',5=>'法检考试',6=>'基层政法机关',7=>'选调生',8=>'事业单位考试',9=>'军转干',10=>'村官考试',11=>'乡镇公务员',12=>'中小学教师招聘',13=>'报关员',14=>'报检员',15=>'司法',16=>'考研',17=>'会计',18=>'路转税',19=>'教师资格证',20=>'党政公选',21=>'特岗教师',22=>'信用社考试',23=>'优职素质教育',24=>'专业课',25=>'银行系统招聘',26=>'三支一扶',27=>'通用考试',28=>'军队文职',29=>'公安边防',30=>'国家电网'') VARCHAR(255)"`
	Area string `json:"area" xorm:"not null default '' comment('所属地区 
999=>'不限',0=>'国家',1=>'北京市',2=>'上海市',3=>'天津市',4=>'重庆市',5=>'河北省',6=>'山西省',7=>'内蒙古',8=>'辽宁省',9=>'吉林省',10=>'黑龙江省',11=>'江苏省',12=>'浙江省',13=>'安徽省',14=>'福建省',15=>'江西省',16=>'山东省',17=>'河南省',18=>'湖北省',19=>'湖南省',20=>'广东省',21=>'广西',22=>'海南省',23=>'四川省',24=>'贵州省',25=>'云南省',26=>'西藏',27=>'陕西省',28=>'甘肃省',29=>'青海省',30=>'宁夏',31=>'新疆',32=>'新疆兵团',101=>'深圳市'') VARCHAR(255)"`
	CourseRelation string    `json:"course_relation" xorm:"not null default '' comment('关联课程') VARCHAR(2550)"`
	Expire         time.Time `json:"expire" xorm:"not null comment('截止日期 格式0000-00-00') DATE"`
	Title          string    `json:"title" xorm:"not null default '' VARCHAR(255)"`
	Content        string    `json:"content" xorm:"not null TEXT"`
	Createdate     time.Time `json:"createdate" xorm:"not null DATETIME"`
	Status         int       `json:"status" xorm:"not null default 2 comment('状态  1:正常 2:取消 3:删除') TINYINT(1)"`
	Lastdate       time.Time `json:"lastdate" xorm:"not null DATETIME"`
	Flag           int       `json:"flag" xorm:"not null default 0 comment('是否置顶 1:是 0:否') TINYINT(1)"`
	Userid         int       `json:"userid" xorm:"not null INT(10)"`
	Username       string    `json:"username" xorm:"not null default '' VARCHAR(20)"`
	Nickname       string    `json:"nickname" xorm:"not null VARCHAR(255)"`
	NoticeDescribe string    `json:"notice_describe" xorm:"not null comment('公告描述') TEXT"`
}
