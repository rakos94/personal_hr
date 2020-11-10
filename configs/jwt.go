package configs

import (
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwtGo.StandardClaims
}

func SetJwt(e *echo.Echo) *echo.Group {
	jwtGroup := e.Group("/api")

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: SigningMethodExample,
		SigningKey:    []byte(SecretKeyExample),
	}))

	return jwtGroup
}

func CreateJwtToken(email string) (string, error) {
	claims := JwtClaims{
		email,
		jwtGo.StandardClaims{
			Id:        email,
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	rawToken := jwtGo.NewWithClaims(jwtGo.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(SecretKeyExample))
	if err != nil {
		return "", err
	}

	return token, nil
}
