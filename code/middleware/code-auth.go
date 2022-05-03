package middleware

var authHead = `package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)`

// %s generate token is generate by param :
// go-gen init-mw user_id:int,role:string, generateToParam
var authInit = `

var (
	key = os.Getenv("JWT_SECRET")
)

type AuthService interface {
	GenerateToken(%s) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authService struct {}

func NewAuthService() *authService {
	return &authService{}
}`

// go-gen init-mw user_id:int,role:string, generateToParam
var authGenerateToken = `
func (s *authService) GenerateToken(%s) (string, error) {
	claims := jwt.MapClaims{%s}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}`

var authValidateToken = `

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
`
