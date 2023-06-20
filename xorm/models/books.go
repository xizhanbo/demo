package models

import (
	"time"
)

type Books struct {
	Id                  int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	RelatedId           string    `json:"related_id" xorm:"comment('套装图书关联图书id,连接') index VARCHAR(20)"`
	Status              int       `json:"status" xorm:"default 0 comment('0未上线 1上线 2下线 3删除') TINYINT(1)"`
	BookName            string    `json:"book_name" xorm:"not null comment('图书名称') index VARCHAR(300)"`
	BookType            int       `json:"book_type" xorm:"default 0 comment('图书类型 0实体图书') TINYINT(1)"`
	IsSuit              int       `json:"is_suit" xorm:"default 0 comment('0非套装图书 1套装图书') TINYINT(4)"`
	Categoryidssubjects string    `json:"categoryIdsSubjects" xorm:"comment('考试类型和科目（1笔试2面试3纯图书4综合）') TEXT"`
	BookCover           string    `json:"book_cover" xorm:"comment('图书封面多张图片,连接') TEXT"`
	Price               string    `json:"price" xorm:"comment('图书原价（元）') index DECIMAL(10,2)"`
	ActualPrice         string    `json:"actual_price" xorm:"comment('真实价格（元）') index DECIMAL(10,2)"`
	SingleCost          string    `json:"single_cost" xorm:"comment('单本书成本价（元）') index DECIMAL(10,2)"`
	Press               string    `json:"press" xorm:"comment('出版社') VARCHAR(200)"`
	BookNo              string    `json:"book_no" xorm:"comment('书号或自编码') index VARCHAR(150)"`
	TotalPage           int       `json:"total_page" xorm:"default 0 comment('总页数') INT(11)"`
	BookPrint           string    `json:"book_print" xorm:"comment('图书印张数') DECIMAL(10,3)"`
	Weight              string    `json:"weight" xorm:"comment('重量kg') VARCHAR(50)"`
	Summary             string    `json:"summary" xorm:"comment('简介') TEXT"`
	CreatorName         string    `json:"creator_name" xorm:"not null comment('创建人用户名') index VARCHAR(100)"`
	CreatedAt           time.Time `json:"created_at" xorm:"index DATETIME"`
	UpdatedAt           time.Time `json:"updated_at" xorm:"DATETIME"`
}
