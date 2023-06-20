package models

import (
	"time"
)

type ProtocolStudent struct {
	Id               int       `json:"id" xorm:"not null pk autoincr comment('当前ID即为新协议的ID') INT(11)"`
	UserId           int64     `json:"user_id" xorm:"not null default 0 comment('用户id,cl_user表userID') index BIGINT(20)"`
	UserName         string    `json:"user_name" xorm:"not null default '' comment('用户名cl_user 表的userName') VARCHAR(50)"`
	ParamId          int       `json:"param_id" xorm:"not null default 0 comment('协议参数id，protocol_parameter表id') INT(11)"`
	TemplateId       int       `json:"template_id" xorm:"not null default 0 comment('协议模板id，protocol_template表id') INT(11)"`
	ContractId       string    `json:"contract_id" xorm:"not null default '' comment('契约锁合同id') VARCHAR(50)"`
	ClassNo          string    `json:"class_no" xorm:"not null default '' comment('课程编号') index VARCHAR(20)"`
	ClassId          int       `json:"class_id" xorm:"not null default 0 comment('课程ID') INT(11)"`
	ClassTitle       string    `json:"class_title" xorm:"not null default '' comment('课程名称') VARCHAR(255)"`
	StudentName      string    `json:"student_name" xorm:"not null default '' comment('学生姓名') VARCHAR(50)"`
	IdCard           string    `json:"id_card" xorm:"not null default '' comment('身份证编号') VARCHAR(50)"`
	Gender           string    `json:"gender" xorm:"not null default '' comment('性别') VARCHAR(5)"`
	Mobile           string    `json:"mobile" xorm:"not null default '' comment('手机号') VARCHAR(12)"`
	Province         string    `json:"province" xorm:"not null default '' comment('考试省份') VARCHAR(20)"`
	Address          string    `json:"address" xorm:"not null default '' comment('地址') VARCHAR(255)"`
	Age              int       `json:"age" xorm:"not null default 0 comment('年龄') TINYINT(2)"`
	TreatyType       int       `json:"treaty_type" xorm:"not null default 0 comment('1笔试，2面试') TINYINT(1)"`
	PayType          int       `json:"pay_type" xorm:"not null default 0 comment('1全款，2预收') TINYINT(1)"`
	Rank             int       `json:"rank" xorm:"not null default 0 comment('名次') INT(10)"`
	MiniScore        float64   `json:"mini_score" xorm:"not null default 0.00 comment('岗位最低分数线') DOUBLE(10,2)"`
	ScoreDiff        int       `json:"score_diff" xorm:"not null default 0 comment('分数差') INT(10)"`
	RecruitsNum      int       `json:"recruits_num" xorm:"not null default 0 comment('岗位招聘人数') INT(10)"`
	ExamStandard     int       `json:"exam_standard" xorm:"not null default 0 comment('补充协议选择标准，当protococl_parameter表pay_type等于2此字段必填， 1代表名次符合状元标准 2代表分数符合状元标准') TINYINT(1)"`
	ExamCardCodePic  string    `json:"exam_card_code_pic" xorm:"not null default '' comment('准考证照片') VARCHAR(255)"`
	ScorePic         string    `json:"score_pic" xorm:"not null default '' comment('成绩单照片') VARCHAR(255)"`
	ExamScore        float64   `json:"exam_score" xorm:"not null comment('成绩分数') DOUBLE(10,2)"`
	ClassPrice       float64   `json:"class_price" xorm:"not null default 0.00 comment('课程价格') DOUBLE(10,2)"`
	ExamDepartment   string    `json:"exam_department" xorm:"not null default '' comment('报考部门') VARCHAR(100)"`
	ExamPosition     string    `json:"exam_position" xorm:"not null default '' comment('报考职位') VARCHAR(100)"`
	ExamPositionCode string    `json:"exam_position_code" xorm:"not null default '' comment('报考职位代码') VARCHAR(100)"`
	ExamCardCode     string    `json:"exam_card_code" xorm:"not null default '' comment('准考证号') VARCHAR(100)"`
	OrderNo          string    `json:"order_no" xorm:"not null default '' comment('订单号cl_order表OrderNum') index VARCHAR(50)"`
	OrderId          int64     `json:"order_id" xorm:"not null comment('cl_order表OrderID') index BIGINT(20)"`
	Status           int       `json:"status" xorm:"not null default 1 comment('协议状态   0失效 1正常') TINYINT(4)"`
	IsCanEdit        int       `json:"is_can_edit" xorm:"not null default 1 comment('1可以编辑，2不能编辑') TINYINT(4)"`
	OperatName       string    `json:"operat_name" xorm:"not null default '' comment('操作人') VARCHAR(50)"`
	UpdateAt         time.Time `json:"update_at" xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
	CreateAt         time.Time `json:"create_at" xorm:"not null default CURRENT_TIMESTAMP comment('添加时间') DATETIME"`
	OldContractId    string    `json:"old_contract_id" xorm:"not null default '' comment('旧合同ID；前端发起重签合同时，将contract_id更新到当前字段') VARCHAR(50)"`
}
