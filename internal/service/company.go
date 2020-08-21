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

func (s *Service) boolQ() *elastic.BoolQuery {
	boolQ := elastic.NewBoolQuery()
	boolQ.Should(elastic.NewTermQuery("regStatus", "存续"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "在业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "在营企业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "开业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "正常"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "在营"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "正常登记"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "登记成立"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "登记"))
	boolQ.Should(elastic.NewTermQuery("regStatus", " 正常"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "正常执业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "已开业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "正常在业"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "核准许可登记"))
	boolQ.Should(elastic.NewTermQuery("regStatus", "正常营业"))
	boolQ.MinimumNumberShouldMatch(1)
	return boolQ
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
	boolQ := s.boolQ()
	if len(req.Name) > 0 {
		query := elastic.NewMatchPhrasePrefixQuery("name", req.Name)
		boolQ.Must(query)
		searchRes, err = s.es.Search(model.ComapnyEsIndex).
			TrackTotalHits(true).
			Query(boolQ).
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

func (s *Service)esSearch(query elastic.Query) (*elastic.SearchResult, error) {
	var searchRes *elastic.SearchResult
	var err error
	searchRes, err = s.es.Search(model.ComapnyEsIndex).
		TrackTotalHits(true).
		Query(query).
		Size(model.DefaultEsSize).
		Do(context.Background())
	if err != nil {
		return nil, err
	}
	if searchRes == nil || searchRes.Hits == nil ||
		searchRes.Hits.TotalHits == nil ||
		searchRes.Hits.TotalHits.Value == 0 {
		return nil, nil
	}
	return searchRes, nil
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

	// 根据法人名称查询一次
	queryName := elastic.NewTermQuery("legalPersonName", req.Name)
	searchRes, err = s.esSearch(queryName)
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	if searchRes == nil {
		// 根据职员名称再查询一次
		queryStaff := elastic.NewTermQuery("staff.name", req.Name)
		nestedQuery := elastic.NewNestedQuery("staff", queryStaff)
		searchRes, err = s.esSearch(nestedQuery)
		if err != nil {
			res.Message = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
	}
	if searchRes == nil {
		res.Message = "No result found"
		c.JSON(http.StatusOK, res)
		return
	}

	itemMap := make(map[string]model.LegalInfo)
	cidsMap := make(map[string][]int64)
	for _, hit := range searchRes.Hits.Hits {
		var compInfo model.CompanyInfo
		err := json.Unmarshal(hit.Source, &compInfo)
		if err != nil {
			s.runLog.Warningf("unmarshal error %s", err.Error())
			continue
		}
		var cids string
		// 根据office里的cid集合做去重
		for _, office := range compInfo.LegalInfo.Office {
			cids += fmt.Sprintf("%d,", office.Cid)
		}
		if _, ok := cidsMap[cids]; !ok {
			cidsMap[cids] = append(cidsMap[cids], compInfo.CompanyId)
			itemMap[cids] = compInfo.LegalInfo
		} else {
			cidsMap[cids] = append(cidsMap[cids], compInfo.CompanyId)
		}
	}

	// 转成列表结构, 并给companyIds赋值
	var items []model.LegalInfo
	for cids, item := range itemMap {
		item.CompanyIds = cidsMap[cids]
		items = append(items, item)
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