package usecase

import (
	"Assignment/pkg/domain"
	"Assignment/pkg/helper"
	interfaces "Assignment/pkg/repository/interface"
	services "Assignment/pkg/usecase/interface"
	"Assignment/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type Adminusecase struct {
	adminrepository interfaces.AdminRepo
}

func NewAdminUsecase(repo interfaces.AdminRepo) services.AdminUsecase {
	return &Adminusecase{
		adminrepository: repo,
	}
}

func (au *Adminusecase) LoginHandler(adminDetails models.Adminlogin) (domain.AdminToken, error) {
	admindetails, err := au.adminrepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.AdminToken{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admindetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.AdminToken{}, err
	}
	var adminDetailResponse models.AdminDetailResponse
	err = copier.Copy(&adminDetailResponse, &admindetails)
	if err != nil {
		return domain.AdminToken{}, err
	}
	token, refresh, err := helper.GenerateAdminToken(adminDetailResponse)
	if err != nil {
		return domain.AdminToken{}, err
	}
	return domain.AdminToken{
		Admin:   adminDetailResponse,
		Token:   token,
		Refresh: refresh,
	}, nil
}
