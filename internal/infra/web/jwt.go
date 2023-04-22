package web

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	secret string
}

type JWTPayload struct {
	Subject    string
	Expiration int64
}

type JWTInterface interface {
	GenerateToken(subject string, expiration time.Duration) (string, error)
	VerifyToken(token string) (*JWTPayload, error)
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) GenerateToken(subject string, expirationSeconds int64) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":    subject,
		"expiration": expirationSeconds,
	})

	return t.SignedString([]byte(j.secret))
}

func (j *JWT) VerifyToken(token string) (*JWTPayload, error) {
	jwtToken, err := jwt.ParseWithClaims(
		token,
		&JWTPayload{},
		func(t *jwt.Token) (interface{}, error) {
			alg := t.Method.Alg()

			if jwt.SigningMethodHS256.Alg() != alg {
				return nil, errors.New("invalid token")
			}
			return []byte(j.secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)

	if !ok {
		return nil, errors.New("invalid token")
	}

	return payload, nil
}

func (p *JWTPayload) Valid() error {

	if time.Now().After(time.Unix(p.Expiration, 0)) {
		return errors.New("token has expired")
	}

	return nil
}
