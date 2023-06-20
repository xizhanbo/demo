package models

import (
	"time"
)

type StarMapLeads struct {
	Id                   int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	LeadsId              int       `json:"leads_id" xorm:"comment('线索id') index INT(11)"`
	TenantId             int       `json:"tenant_id" xorm:"comment('商户id') INT(11)"`
	MsgId                string    `json:"msg_id" xorm:"comment('消息id') unique VARCHAR(50)"`
	OpportunityId        int       `json:"opportunity_id" xorm:"comment('机会id') INT(11)"`
	OpportunityPoolCode  string    `json:"opportunity_pool_code" xorm:"default '' comment('线索机会池code，不传则为默认创建的线索机会池') VARCHAR(200)"`
	LeadsSource          string    `json:"leads_source" xorm:"default '' comment('线索来源') VARCHAR(150)"`
	Operator             string    `json:"operator" xorm:"default '' comment('创建人') VARCHAR(70)"`
	ChatId               string    `json:"chat_id" xorm:"default '' comment('线索聊天会话唯一标识') VARCHAR(150)"`
	Mobile               string    `json:"mobile" xorm:"default '' comment('手机号') index VARCHAR(20)"`
	Mobilebak            string    `json:"mobilebak" xorm:"default '' comment('备用手机号') index VARCHAR(20)"`
	Phone                string    `json:"phone" xorm:"default '' comment('座机号') VARCHAR(20)"`
	Wechat               string    `json:"wechat" xorm:"default '' comment('微信号') VARCHAR(50)"`
	Qq                   string    `json:"qq" xorm:"default '' comment('qq号') VARCHAR(20)"`
	Mail                 string    `json:"mail" xorm:"default '' comment('邮箱') VARCHAR(150)"`
	UserName             string    `json:"user_name" xorm:"default '' comment('学员用户名') VARCHAR(100)"`
	Sex                  int       `json:"sex" xorm:"default 0 comment('学员性别：1-男，2-女') TINYINT(1)"`
	IpCountryCode        string    `json:"ip_country_code" xorm:"default '' comment('ip国家code') VARCHAR(20)"`
	IpProvinceCode       string    `json:"ip_province_code" xorm:"default '' comment('ip省份code') VARCHAR(20)"`
	IpCityCode           string    `json:"ip_city_code" xorm:"default '' comment('ip城市code') VARCHAR(20)"`
	ProvinceCode         string    `json:"province_code" xorm:"default '' comment('推广省份code') VARCHAR(20)"`
	CityCode             string    `json:"city_code" xorm:"default '' comment('推广城市code') VARCHAR(20)"`
	SkuTypeCode          string    `json:"sku_type_code" xorm:"default '' comment('sku分类') VARCHAR(20)"`
	SkuCode              string    `json:"sku_code" xorm:"default '' comment('sku') VARCHAR(255)"`
	Pvid                 string    `json:"pvid" xorm:"default '' comment('在线索获取过程中获得的用户账号') VARCHAR(160)"`
	Uid                  string    `json:"uid" xorm:"default '' comment('在线索获取过程中获得的用户账号') VARCHAR(150)"`
	Uuid                 string    `json:"uuid" xorm:"default '' comment('同一用户，尚德内部唯一标识') VARCHAR(150)"`
	Openid               string    `json:"openid" xorm:"default '' comment('用户在微信端的id信息，可用于微信精灵、小程序等端的使用') VARCHAR(150)"`
	MediaAttributions    string    `json:"media_attributions" xorm:"default '' comment('线索在投放端的媒体属性') VARCHAR(2000)"`
	BusinessAttributions string    `json:"business_attributions" xorm:"default '' comment('线索在投放端配置的业务属性') VARCHAR(2000)"`
	PromoterDepartment   string    `json:"promoter_department" xorm:"default '' comment('线索对应推广人的组织层级结构') VARCHAR(2000)"`
	ReceiverDepartment   string    `json:"receiver_department" xorm:"default '' comment('线索对应接收方的组织层级结构') VARCHAR(2000)"`
	Url                  string    `json:"url" xorm:"default '' comment('线索录入的落地页地址') VARCHAR(2500)"`
	Ip                   string    `json:"ip" xorm:"default '' comment('ip') VARCHAR(50)"`
	DeviceCode           string    `json:"device_code" xorm:"default '' comment('设备类型') VARCHAR(50)"`
	SearchWord           string    `json:"search_word" xorm:"default '' comment('搜索词') VARCHAR(2000)"`
	LeadsUnitCost        string    `json:"leads_unit_cost" xorm:"default '' comment('单个线索的投放/获取成本') VARCHAR(20)"`
	CustomFields         string    `json:"custom_fields" xorm:"default '' VARCHAR(2000)"`
	PlatformName         string    `json:"platform_name" xorm:"comment('平台名称') VARCHAR(200)"`
	AdPlanId             string    `json:"ad_plan_id" xorm:"comment('广告计划id') index VARCHAR(100)"`
	AdPlanName           string    `json:"ad_plan_name" xorm:"comment('广告计划名称') index VARCHAR(300)"`
	AdGroupId            string    `json:"ad_group_id" xorm:"comment('广告组id') index VARCHAR(100)"`
	AdGroupName          string    `json:"ad_group_name" xorm:"comment('广告组名称') index VARCHAR(300)"`
	AdIdeaId             string    `json:"ad_idea_id" xorm:"comment('广告创意id') index VARCHAR(100)"`
	AdIdeaName           string    `json:"ad_idea_name" xorm:"comment('广告创意名称') index VARCHAR(300)"`
	KeyWordId            string    `json:"key_word_id" xorm:"comment('关键词id') VARCHAR(100)"`
	KeyWordName          string    `json:"key_word_name" xorm:"comment('关键词名称') VARCHAR(300)"`
	FetchTime            time.Time `json:"fetch_time" xorm:"comment('线索获取时间') DATETIME"`
	UpdatedAt            time.Time `json:"updated_at" xorm:"DATETIME"`
	CreatedAt            time.Time `json:"created_at" xorm:"index DATETIME"`
}
