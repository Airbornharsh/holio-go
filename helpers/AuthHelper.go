package helpers

import (
	"os"
	"regexp"

	"github.com/airbornharsh/holio-go/models"
	"github.com/dgrijalva/jwt-go"
)

func IsValidPassword(password string) bool {
	passwordPatternRegex := "[a-zA-Z0-9!@#$%^&*()_+{}|:<>?]{8,}"

	regexpObj := regexp.MustCompile(passwordPatternRegex)

	return regexpObj.MatchString(password)
}

func GenerateToken(user *models.User) (string, error) {
	JWTSECRET := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"phone":    user.PhoneNumber,
		"type":     user.UserType,
		"exp":      15000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWTSECRET))
}
