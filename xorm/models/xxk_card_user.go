package models

import (
	"time"
)

type XxkCardUser struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LoginName  string    `json:"login_name" xorm:"not null default '' comment('登录的用户名，可能是手机号，邮箱，或真实用户名') VARCHAR(20)"`
	Username   string    `json:"username" xorm:"not null default '' comment('用户名') VARCHAR(20)"`
	Nickname   string    `json:"nickname" xorm:"not null default '' comment('供页面显示的代理商名次') VARCHAR(40)"`
	Role       int       `json:"role" xorm:"not null default 2 comment('角色(1：一级代理，2：二级代理，3：大区经理，4：管理员）') TINYINT(4)"`
	Parent     string    `json:"parent" xorm:"not null default '' comment('归属上级的用户名') VARCHAR(20)"`
	AllowLogon int       `json:"allow_logon" xorm:"not null default 0 comment('是否允许登录（0：不允许，1允许）') TINYINT(4)"`
	Operator   string    `json:"operator" xorm:"not null default '' comment('记录添加绑定操作的用户名') VARCHAR(20)"`
	CreateTime time.Time `json:"create_time" xorm:"not null default CURRENT_TIMESTAMP comment('绑定帐号的时间') TIMESTAMP"`
}
