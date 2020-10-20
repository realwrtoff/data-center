package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type NewSeedProject struct {
	// Id              int       `gorm:"type:int(11) auto_increment;primary_key" json:"id"`
	ProjectId       int    `gorm:"type:int(11);unique_index:newseed_project_uniq_idx;not null;default:'0';comment:'项目ID'" json:"project_id"`
	Name            string `gorm:"type:varchar(255);unique_index:newseed_project_uniq_idx;not null;default:'';comment:'项目名称'" json:"name"`
	Company         string `gorm:"type:varchar(50);not null;default:'';comment:'公司名称'" json:"company"`
	Fname           string `gorm:"type:varchar(50);not null;default:'';comment:'分类名称'" json:"fname"`
	Province        string `gorm:"type:varchar(50);not null;default:'';comment:'省份'" json:"province"`
	City            string `gorm:"type:varchar(50);not null;default:'';comment:'城市'" json:"city"`
	EstablishedTime string `gorm:"type:varchar(50);not null;default:'';comment:'成立时间'" json:"established_time"`
	Status          string `gorm:"type:varchar(50);not null;default:'';comment:'营业状态'" json:"status"`
	Desc            string `gorm:"type:text; not null;comment:'简介'" json:"desc"`
	Competitor      string `gorm:"type:text; not null;comment:'竞品'" json:"competitor"`
	Logo            string `gorm:"type:text; not null;comment:'logo';comment:'项目ID'" json:"logo"`
	Tags            string `gorm:"type:text; not null;comment:'标签'" json:"tags"`
	// UpdateTime      time.Time `gorm:"default:null" json:"update_time"`
}


type NewSeedProjectInvest struct {
	ProjectId int    `json:"project_id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Round     string `json:"round"`
	Amt       string `json:"amt"`
	Currency  string `json:"currency"`
	RoundDate string `json:"round_date"`
	Investor  string `json:"investor"`
}

// 自定义表名
func (NewSeedProject) TableName() string {
	return "newseed_project"
}

func (NewSeedProjectInvest) TableName() string {
	return "newseed_project_invest"
}
