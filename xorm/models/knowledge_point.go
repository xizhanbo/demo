package models

type KnowledgePoint struct {
	Pukey            int     `json:"PUKEY" xorm:"not null pk autoincr comment('知识点ID') INT(10)"`
	Name             string  `json:"name" xorm:"not null default '' comment('知识点名称') VARCHAR(30)"`
	Descrp           string  `json:"descrp" xorm:"not null default '' comment('知识点描述') VARCHAR(255)"`
	DifficultGrade   float64 `json:"difficult_grade" xorm:"not null comment('难度系数') DOUBLE"`
	AppearGrade      int     `json:"appear_grade" xorm:"not null default 0 comment('高频系数') TINYINT(1)"`
	ImportantGrade   int     `json:"important_grade" xorm:"not null default 0 comment('重点系数') TINYINT(1)"`
	MistakeGrade     int     `json:"mistake_grade" xorm:"not null default 0 comment('易错系数') TINYINT(1)"`
	DistinctionGrade float64 `json:"distinction_grade" xorm:"not null default 1 comment('区分度') DOUBLE"`
	GuessGrade       float64 `json:"guess_grade" xorm:"not null default 0.2 comment('猜测值') DOUBLE"`
	PrevKp           int     `json:"prev_kp" xorm:"not null default 0 comment('上一级知识点') INT(10)"`
	NodeRank         int     `json:"node_rank" xorm:"not null default 0 comment('节点所在层级') TINYINT(1)"`
	IsLeafnode       int     `json:"is_leafnode" xorm:"not null default 1 comment('是否为叶子节点') TINYINT(1)"`
	IsExtQuestion    int     `json:"is_ext_question" xorm:"not null default 1 comment('是否可以参与抽题') TINYINT(1)"`
	QuestionNum      int     `json:"question_num" xorm:"not null default 0 comment('现有题量') INT(10)"`
	BlSub            int     `json:"bl_sub" xorm:"not null default 0 comment('所属科目') INT(10)"`
	Fb1z1            string  `json:"FB1Z1" xorm:"not null default '' comment('备用1') VARCHAR(100)"`
	Fb1z2            string  `json:"FB1Z2" xorm:"not null default '' comment('备用2') VARCHAR(100)"`
	Fb1z3            string  `json:"FB1Z3" xorm:"not null default '' comment('备用3') VARCHAR(100)"`
	Fb1z4            string  `json:"FB1Z4" xorm:"not null default '' comment('备用4') VARCHAR(100)"`
	Fb1z5            string  `json:"FB1Z5" xorm:"not null default '' comment('备用5') VARCHAR(100)"`
	Eb1b1            string  `json:"EB1B1" xorm:"not null comment('扩展1') VARCHAR(100)"`
	Eb102            string  `json:"EB102" xorm:"not null default '' comment('扩展2') VARCHAR(100)"`
	Eb103            string  `json:"EB103" xorm:"not null default '' comment('扩展3') VARCHAR(100)"`
	Eb104            string  `json:"EB104" xorm:"not null default '' comment('扩展4') VARCHAR(100)"`
	Eb105            string  `json:"EB105" xorm:"not null default '' comment('扩展5') VARCHAR(100)"`
	Bb1b1            int     `json:"BB1B1" xorm:"not null default -1 comment('审核标示') TINYINT(1)"`
	Bb102            int     `json:"BB102" xorm:"not null default 1 comment('有效标识') TINYINT(1)"`
	Bb103            int     `json:"BB103" xorm:"not null default 0 comment('创建时间') INT(10)"`
	Bb104            string  `json:"BB104" xorm:"not null default '' comment('创建者') VARCHAR(20)"`
	Bb105            int     `json:"BB105" xorm:"not null default 0 comment('创建者ID') INT(10)"`
	Bb106            int     `json:"BB106" xorm:"not null default 0 comment('更新时间') INT(10)"`
	Bb107            string  `json:"BB107" xorm:"not null default '' comment('更新者') VARCHAR(20)"`
	Bb108            int     `json:"BB108" xorm:"not null default 0 comment('更新者ID') INT(10)"`
}
