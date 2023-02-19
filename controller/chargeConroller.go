package controller

import (
	"github.com/gin-gonic/gin"
	"soft-pro/resp"
	"soft-pro/service"
)

// POST /user/station/list 查看充电站列表
func StationList(c *gin.Context) {
	css, err := service.GetStationList()
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
	}
	resp.OkWithData(css, c)
}

// GET /user/points/:StationID 查看充电站下所有充电桩
func PointList(c *gin.Context) {
	StationID := c.Param("StationID")
	css, err := service.GetPointList(StationID)
	if err != nil {
		resp.FailWithMessage(err.Error(), c)
	}
	resp.OkWithData(css, c)
}
