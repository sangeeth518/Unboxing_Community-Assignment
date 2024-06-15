package helper

import (
	"Assignment/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateAdminToken(admin models.AdminDetailResponse) (string, string, error) {
	tokenClaimes := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaims := &AuthCustomClaims{
		Id: admin.ID,
		// Name: admin.Name,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaimes)
	tokenString, err := token.SignedString([]byte("adminsecret"))

	// fmt.Println("in admin helper",tokenString)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("adminrefresh"))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}
