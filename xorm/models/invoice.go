package models

import (
	"time"
)

type Invoice struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	OrderId        int       `json:"order_id" xorm:"not null comment('订单id') index INT(11)"`
	OrderNum       string    `json:"order_num" xorm:"not null comment('订单编号') index VARCHAR(150)"`
	OrderDate      time.Time `json:"order_date" xorm:"comment('订单下单时间') DATETIME"`
	Fpqqlsh        string    `json:"fpqqlsh" xorm:"comment('发票请求流水号') VARCHAR(150)"`
	InvoiceId      string    `json:"invoice_id" xorm:"comment('诺诺网发票主键id') VARCHAR(50)"`
	InvoiceMoney   string    `json:"invoice_money" xorm:"comment('发票金额') DECIMAL(10,2)"`
	InvoiceType    int       `json:"invoice_type" xorm:"not null default 0 comment('0电子1纸质') TINYINT(1)"`
	TitleType      int       `json:"title_type" xorm:"default 0 comment('0个人1单位') TINYINT(1)"`
	Status         int       `json:"status" xorm:"not null default 0 comment('0已提交1已开票2发票异常') TINYINT(1)"`
	InvoiceCode    string    `json:"invoice_code" xorm:"comment('发票代码') VARCHAR(200)"`
	InvoiceNum     string    `json:"invoice_num" xorm:"comment('发票号码') VARCHAR(200)"`
	InvoiceTitle   string    `json:"invoice_title" xorm:"not null comment('发票抬头') VARCHAR(200)"`
	TaxNum         string    `json:"tax_num" xorm:"comment('纳税人识别号') VARCHAR(200)"`
	InvoiceContent string    `json:"invoice_content" xorm:"comment('发票内容') VARCHAR(150)"`
	CancelTag      int       `json:"cancel_tag" xorm:"default 0 comment('0未作废1已作废(冲红)') TINYINT(1)"`
	PdfLink        string    `json:"pdf_link" xorm:"comment('发票pdf地址') VARCHAR(255)"`
	Operator       string    `json:"operator" xorm:"comment('管理员用户名') VARCHAR(100)"`
	CreatedAt      time.Time `json:"created_at" xorm:"DATETIME"`
	UpdatedAt      time.Time `json:"updated_at" xorm:"DATETIME"`
}
