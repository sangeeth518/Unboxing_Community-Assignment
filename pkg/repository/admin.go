package repository

import (
	"Assignment/pkg/domain"
	interfaces "Assignment/pkg/repository/interface"
	"Assignment/pkg/utils/models"

	"gorm.io/gorm"
)

type Adminrepo struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepo {
	return &Adminrepo{
		DB: DB,
	}
}

func (ad *Adminrepo) LoginHandler(Admindetails models.Adminlogin) (domain.Admin, error) {
	var admin domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ?", Admindetails.Email).Scan(&admin).Error; err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}
