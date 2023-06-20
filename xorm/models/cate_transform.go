package models

type CateTransform struct {
	Id             int    `json:"id" xorm:"not null pk autoincr comment('id') INT(11)"`
	PhpCateId      int    `json:"php_cate_id" xorm:"not null default 0 comment('php项目配置中的类型ID，例：1000=公务员') unique(u_php_java_cate_id) INT(11)"`
	JavaCateId     int    `json:"java_cate_id" xorm:"not null default 0 comment('java的类型ID，例：1=公务员') unique(u_php_java_cate_id) INT(11)"`
	ChildrenCateId string `json:"children_cate_id" xorm:"comment('php 分类id') VARCHAR(255)"`
}
