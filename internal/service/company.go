package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/realwrtoff/data-center/internal/model"
	"math"
	"net/http"
)

// 根据公司名称或者老板名称查询
type NameSearchReq struct {
	Name string `form:"name" json:"name,omitempty"`
	Page int    `form:"page" json:"page,omitempty"`
	Size int    `form:"size" json:"size,omitempty"`
}

// 根据公司ID查询
type CompanyIdReq struct {
	CompanyId string `form:"company_id" json:"company_id,omitempty"`
}

func (s *Service) CompanySearch(c *gin.Context) {
	req := &NameSearchReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}

	res := &BasicRes{
		Code:    200,
		Message: "success",
	}

	var searchRes *elastic.SearchResult
	var err error
	if len(req.Name) > 0 {
		query := elastic.NewMatchQuery("name", req.Name)
		searchRes, err = s.es.Search(model.ComapnyEsIndex).
			TrackTotalHits(true).
			Query(query).
			Size(req.Size).
			From((req.Page - 1) * req.Size).
			Do(context.Background())
	} else {
		searchRes, err = s.es.Search(model.ComapnyEsIndex).
			TrackTotalHits(true).
			Size(req.Size).
			From((req.Page - 1) * req.Size).
			Do(context.Background())
	}

	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if searchRes == nil || searchRes.Hits == nil ||
		searchRes.Hits.TotalHits == nil ||
		searchRes.Hits.TotalHits.Value == 0 {
		res.Message = "No result found"
		c.JSON(http.StatusOK, res)
		return
	}

	var items []model.CompanyInfo
	for _, hit := range searchRes.Hits.Hits {
		var compInfo model.CompanyInfo
		err := json.Unmarshal(hit.Source, &compInfo)
		if err != nil {
			s.runLog.Warningf("unmarshal error %s", err.Error())
			continue
		}
		items = append(items, compInfo)
	}
	totalCount := int(searchRes.Hits.TotalHits.Value)
	// 第一次感受到go如此蛋疼，一个向上取整而已，艹蛋
	totalPage := int(math.Ceil(float64(totalCount / req.Size)))

	pageDataRes := &PageDataRes{
		Page:       req.Page,
		TotalCount: totalCount,
		TotalPage:  totalPage,
		Items:      items,
	}
	res.Data = pageDataRes

	c.JSON(http.StatusOK, res)
}

func (s *Service) CompanyDetail(c *gin.Context) {
	req := &CompanyIdReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	query := elastic.NewTermQuery("companyId", req.CompanyId)
	searchRes, err := s.es.Search(model.ComapnyEsIndex).Query(query).Do(context.Background())
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if searchRes == nil || searchRes.Hits == nil ||
		searchRes.Hits.TotalHits == nil ||
		searchRes.Hits.TotalHits.Value == 0 {
		res.Message = "No result found"
		c.JSON(http.StatusOK, res)
		return
	}

	// var items []QccCompanyInfo
	for _, hit := range searchRes.Hits.Hits {
		var compInfo model.CompanyInfo
		err := json.Unmarshal(hit.Source, &compInfo)
		if err != nil {
			s.runLog.Warningf("unmarshal error %s", err.Error())
			continue
		}
		// items = append(items, compInfo)
		res.Data = compInfo
		break
	}
	c.JSON(http.StatusOK, res)
}

func (s *Service) PersonSearch(c *gin.Context) {
	req := &NameSearchReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	res := &BasicRes{
		Code:    200,
		Message: "success",
	}

	var searchRes *elastic.SearchResult
	var err error

	queryName := elastic.NewMatchQuery("legalInfo.name", req.Name)
	query := elastic.NewNestedQuery("legalInfo", queryName)
	searchRes, err = s.es.Search(model.ComapnyEsIndex).
		TrackTotalHits(true).
		Query(query).
		Size(model.DefaultEsSize).
		Do(context.Background())

	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if searchRes == nil || searchRes.Hits == nil ||
		searchRes.Hits.TotalHits == nil ||
		searchRes.Hits.TotalHits.Value == 0 {
		res.Message = "No result found"
		c.JSON(http.StatusOK, res)
		return
	}

	var items []model.LegalInfo
	cidsMap := make(map[string]int)
	for _, hit := range searchRes.Hits.Hits {
		var compInfo model.CompanyInfo
		err := json.Unmarshal(hit.Source, &compInfo)
		if err != nil {
			s.runLog.Warningf("unmarshal error %s", err.Error())
			continue
		}
		var cids string
		for _, office := range compInfo.LegalInfo.Office {
			cids += fmt.Sprintf("%d,", office.Cid)
		}
		if _, ok := cidsMap[cids]; !ok {
			cidsMap[cids] = 1
			items = append(items, compInfo.LegalInfo)
		}
	}

	pageDataRes := &PageDataRes{
		Page:       req.Page,
		TotalCount: len(items),
		TotalPage:  1,
		Items:      items,
	}
	res.Data = pageDataRes

	c.JSON(http.StatusOK, res)
}