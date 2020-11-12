package configs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func SetJwt(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/api")

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: SigningMethodExample,
		SigningKey:    []byte(SecretKeyExample),
	}))

	return jwtGroup
}

func CreateJwtToken(email string) (string, error) {
	// Create token
	var signingMethod *jwt.SigningMethodHMAC
	if SigningMethodExample == "HS512" {
		signingMethod = jwt.SigningMethodHS512
	}
	token := jwt.New(signingMethod)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Duration(JwtExpiredHour) * time.Hour).Unix()

	t, err := token.SignedString([]byte(SecretKeyExample))
	if err != nil {
		return "", err
	}

	return t, nil
}
