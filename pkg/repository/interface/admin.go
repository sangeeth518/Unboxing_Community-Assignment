package interfaces

import (
	"Assignment/pkg/domain"
	"Assignment/pkg/utils/models"
)

type AdminRepo interface {
	LoginHandler(Admindetails models.Adminlogin) (domain.Admin, error)
}
