package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/realwrtoff/data-center/internal/model"
	"net/http"
)

// 根据公司名称或者老板名称查询
type ProjectReq struct {
	ProjectId int `form:"project_id" json:"project_id,omitempty"`
	Name string `form:"name" json:"name,omitempty"`
	Page int    `form:"page" json:"page,omitempty"`
	Size int    `form:"size" json:"size,omitempty"`
}

func (s *Service) ProjectSearch(c *gin.Context) {
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
	var projects []model.NewSeedProject
	if req.Name != "" {
		word := fmt.Sprintf("%%%s%%", req.Name)
		err = s.mdb.Where("name like ?", word).Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	} else {
		err = s.mdb.Limit(req.Size).Offset((req.Page-1)*req.Size).Find(&projects).Error
	}
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Data = projects
	c.JSON(http.StatusOK, res)
}

func (s *Service) ProjectDetail(c *gin.Context) {
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
	var project model.NewSeedProject
	err = s.mdb.Where("project_id = ?", req.ProjectId).Find(&project).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Data = &project
	c.JSON(http.StatusOK, res)
}

func (s *Service) ProjectInvest(c *gin.Context) {
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
	var invests []model.NewSeedProjectInvest
	err = s.mdb.Where("project_id = ?", req.ProjectId).Find(&invests).Error
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
	}
	res.Data = invests
	c.JSON(http.StatusOK, res)
}

