package handler

import (
	interfaces "Assignment/pkg/usecase/interface"
	"Assignment/pkg/utils/models"
	"Assignment/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminaHandler struct {
	adminusecase interfaces.AdminUsecase
}

func NewAdminHandler(usecase interfaces.AdminUsecase) *AdminaHandler {
	return &AdminaHandler{
		adminusecase: usecase,
	}
}

func (ad *AdminaHandler) LoginHandler(c *gin.Context) {
	var admindetails models.Adminlogin
	if err := c.BindJSON(&admindetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ad.adminusecase.LoginHandler(admindetails)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate admin", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetCookie("Authorization", admin.Token, 3600*24*30, "", "", false, true)
	c.SetCookie("Refresh", admin.Refresh, 3600*24*30, "", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "Admin logged in succesfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}
