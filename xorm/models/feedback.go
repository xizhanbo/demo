package models

type Feedback struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Feedbackcategoryid int    `json:"FeedbackCategoryId" xorm:"not null comment('反馈分类ID') INT(10)"`
	Pagechooseid       int    `json:"PageChooseId" xorm:"not null comment('页面选择项ID') INT(10)"`
	Feeddetail         string `json:"FeedDetail" xorm:"not null comment('反馈具体内容') VARCHAR(255)"`
	Createdate         int    `json:"CreateDate" xorm:"not null comment('反馈时间') INT(10)"`
	Userid             int    `json:"UserId" xorm:"not null comment('用户ID') INT(10)"`
	Username           string `json:"UserName" xorm:"not null comment('用户名') VARCHAR(50)"`
	Replycontent       string `json:"ReplyContent" xorm:"comment('网校老师的回复（允许空值）') VARCHAR(255)"`
	Replydate          int    `json:"ReplyDate" xorm:"comment('回复时间(允许空值)') INT(10)"`
	Status             int    `json:"status" xorm:"default 0 SMALLINT(2)"`
	Lastreplyuserid    int    `json:"LastReplyUserId" xorm:"INT(11)"`
	Replynum           int    `json:"ReplyNum" xorm:"default 0 comment('回复数') INT(11)"`
	Lastreplyusername  string `json:"LastReplyUserName" xorm:"VARCHAR(50)"`
}
