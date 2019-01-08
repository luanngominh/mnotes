package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luanngominh/mnotes/backend/config"
)

const (
	hoursInDay       = 24
	minutesInHour    = 60
	dayInYear        = 365
	miniuteInOneYear = hoursInDay * minutesInHour * dayInYear
)

//JWTPayload ...
type JWTPayload struct {
	*jwt.StandardClaims
	*AuthPayload
}

//AuthPayload user info which is stored in Jwt token
type AuthPayload struct {
	ID    string `json:"id"`
	Name  string `json:"namme"`
	Email string `json:"email"`
}

//GenerateJWTToken generate jwt token for manage user session
func GenerateJWTToken(id, name, email string) (string, error) {
	expires := time.Now().Add(time.Minute * miniuteInOneYear).Unix()

	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = &JWTPayload{
		&jwt.StandardClaims{
			ExpiresAt: expires,
		},
		&AuthPayload{
			Name:  name,
			ID:    id,
			Email: email,
		},
	}

	tokenString, err := t.SignedString(config.Cfg.SignKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
