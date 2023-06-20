package models

import (
	"time"
)

type TestExercises struct {
	Id               int       `json:"id" xorm:"not null pk INT(11)"`
	Paperid          int       `json:"paperid" xorm:"not null INT(11)"`
	KnowledgePointid string    `json:"knowledge_pointid" xorm:"default '' VARCHAR(255)"`
	Subjectid        int       `json:"subjectid" xorm:"TINYINT(4)"`
	Question         string    `json:"question" xorm:"not null VARCHAR(4000)"`
	Choice1          string    `json:"choice1" xorm:"VARCHAR(255)"`
	Choice2          string    `json:"choice2" xorm:"VARCHAR(255)"`
	Choice3          string    `json:"choice3" xorm:"VARCHAR(255)"`
	Choice4          string    `json:"choice4" xorm:"VARCHAR(255)"`
	Choice5          string    `json:"choice5" xorm:"VARCHAR(255)"`
	Choice6          string    `json:"choice6" xorm:"VARCHAR(255)"`
	Choice7          string    `json:"choice7" xorm:"VARCHAR(255)"`
	Testno           int       `json:"testno" xorm:"TINYINT(4)"`
	Testtype         int       `json:"testtype" xorm:"not null TINYINT(4)"`
	Answer           string    `json:"answer" xorm:"VARCHAR(255)"`
	Skill            string    `json:"skill" xorm:"VARCHAR(255)"`
	Expand           string    `json:"expand" xorm:"VARCHAR(255)"`
	Paraphrase       string    `json:"paraphrase" xorm:"VARCHAR(255)"`
	ErrorCorrection  string    `json:"error_correction" xorm:"VARCHAR(255)"`
	Analysis         string    `json:"analysis" xorm:"VARCHAR(4000)"`
	Createdate       time.Time `json:"createdate" xorm:"not null DATETIME"`
	Difficult        float32   `json:"difficult" xorm:"default 0 FLOAT"`
	Score            float32   `json:"score" xorm:"default 0 FLOAT"`
	Writer           string    `json:"writer" xorm:"VARCHAR(50)"`
	Source           string    `json:"source" xorm:"VARCHAR(255)"`
	Verifier         string    `json:"verifier" xorm:"VARCHAR(255)"`
	Status           int       `json:"status" xorm:"not null default 1 TINYINT(1)"`
}
