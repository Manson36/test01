package repositories

import "github.com/test/adminUser/datamodels"

type AdminUserRepository interface {
	Create(user *datamodels.AdminUser) error
}
