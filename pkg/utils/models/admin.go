package models

type Adminlogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AdminDetailResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
