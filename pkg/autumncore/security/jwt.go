package security

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JwtAuth struct {
	JwtKey string
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{}
}

func (j *JwtAuth) SetJwtKey(key string) *JwtAuth {
	j.JwtKey = key
	return j
}

func (j *JwtAuth) CreateToken(claims jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedInToken, err := token.SignedString([]byte(j.JwtKey))
	if err != nil {
		return "", fmt.Errorf("error when creating token JWT %v", err.Error())
	}

	return SignedInToken, nil
}

func (j *JwtAuth) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unable create jwt token")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(j.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Invalid token string ")
	}
}
