package domain

import "Assignment/pkg/utils/models"

type Admin struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminToken struct {
	Admin   models.AdminDetailResponse
	Token   string
	Refresh string
}
