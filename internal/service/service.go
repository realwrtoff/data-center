package service

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type Service struct {
	rds    *redis.Client
	mdb    *gorm.DB
	es     *elastic.Client
	runLog *logrus.Logger
}

func NewService(
	rds *redis.Client,
	mdb *gorm.DB,
	es *elastic.Client,
	runLog *logrus.Logger,
) *Service {
	return &Service{
		rds:    rds,
		mdb:    mdb,
		es:     es,
		runLog: runLog,
	}
}

type BasicRes struct {
	Code    int         `form:"code" json:"code"`
	Message string      `form:"message" json:"message"`
	Data    interface{} `form:"data" json:"data"`
}

type PageDataRes struct {
	Page       int         `json:"page"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}
