package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/test/adminUser/models"
	"github.com/test/adminUser/services"
	"net/http"
)

type AdminUserController struct {
	service services.AdminUserService
}

func (a AdminUserController) Create(c *gin.Context) {
	body := models.AdminUserReqBody{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusOK, models.Ret{Code: 400, Msg: "创建管理用户解析失败", Data: nil})
		return
	}

	c.JSON(http.StatusOK, a.service.Create(&body))
}
