package interfaces

import (
	"Assignment/pkg/domain"
	"Assignment/pkg/utils/models"
)

type AdminUsecase interface {
	LoginHandler(adminDetails models.Adminlogin) (domain.AdminToken, error)
}
