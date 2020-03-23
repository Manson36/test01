package services

import (
	"github.com/test/adminUser/models"
	"github.com/test/adminUser/repositories"
	"log"
	"strings"
)

type AdminUserService interface {
	Create(body *models.AdminUserReqBody) *models.Ret
}

type adminUserService struct {
	repo repositories.AdminUserRepository
}

func (a adminUserService) Create(body *models.AdminUserReqBody) *models.Ret {
	id := 10001

	nickname := strings.TrimSpace(body.Nickname)
	if nickname == "" {
		return &models.Ret{Code: 400, Msg: "请输入用户昵称"}
	}

	password := strings.TrimSpace(body.Password)
	if password == "" {
		return &models.Ret{Code: 400, Msg: "请输入用户密码"}
	}

	email := strings.TrimSpace(body.Email)
	if email == "" {
		return &models.Ret{Code: 400, Msg: "请输入用户邮箱"}
	}

}
