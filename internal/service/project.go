package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realwrtoff/data-center/internal/model/project"
	"net/http"
)

// 根据公司名称或者老板名称查询
type ProjectReq struct {
	ProjectId int `form:"project_id" json:"project_id,omitempty"`
	Name string `form:"name" json:"name,omitempty"`
	Page int    `form:"page" json:"page,omitempty"`
	Size int    `form:"size" json:"size,omitempty"`
}

// 根据ID查询发布人
type PublisherReq struct {
	PublisherId int `form:"publisher_id" json:"publisher_id"`
}

func (s *Service) NewSeedSearch(c *gin.Context) {
	req := &ProjectReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	var err error
	var projects []project.NewSeedProject
	if req.Name != "" {
		word := fmt.Sprintf("%%%s%%", req.Name)
		err = s.mdb.Where("name like ?", word).Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	} else {
		err = s.mdb.Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	}
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Data = projects
	c.JSON(http.StatusOK, res)
}

func (s *Service) NewSeedDetail(c *gin.Context) {
	req := &ProjectReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	var err error
	var newSeedProject project.NewSeedProject
	err = s.mdb.Where("project_id = ?", req.ProjectId).Find(&newSeedProject).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Data = &newSeedProject
	c.JSON(http.StatusOK, res)
}

func (s *Service) NewSeedInvest(c *gin.Context) {
	req := &ProjectReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	var err error
	var invests []project.NewSeedProjectInvest
	err = s.mdb.Where("project_id = ?", req.ProjectId).Find(&invests).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Data = invests
	c.JSON(http.StatusOK, res)
}

func (s *Service) AiHeHuoSearch(c *gin.Context) {
	req := &ProjectReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.Size == 0 {
		req.Size = 10
	}
	var err error
	var projects []project.AiHeHuoProject
	if req.Name != "" {
		word := fmt.Sprintf("%%%s%%", req.Name)
		err = s.mdb.Where("name like ?", word).Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	} else {
		err = s.mdb.Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	}
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	for i := 0; i < len(projects); i++ {
		projects[i].JsonDeal()
	}
	res.Data = projects
	c.JSON(http.StatusOK, res)
}

func (s *Service) AiHeHuoDetail(c *gin.Context) {
	req := &ProjectReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	var err error
	var prj project.AiHeHuoProject
	err = s.mdb.Where("project_id = ?", req.ProjectId).Find(&prj).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		s.runLog.Errorf("find failed [%s]", err.Error())
		return
	}
	prj.JsonDeal()
	res.Data = prj
	c.JSON(http.StatusOK, res)
}

func (s *Service) AiHeHuoPublisher(c *gin.Context) {
	req := &PublisherReq{}

	if err := c.Bind(req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	res := &BasicRes{
		Code:    200,
		Message: "success",
	}
	var err error
	var publisher project.AiHeHuoPublisher
	err = s.mdb.Where("publisher_id = ?", req.PublisherId).Find(&publisher).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	publisher.JsonDeal()
	res.Data = publisher
	c.JSON(http.StatusOK, res)
}

