package jwt

import (
	"errors"
	"pkg/static"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(dependencies JWTDependencies) *JWT {
	expiresIn := time.Duration(dependencies.Config.Security.JWT.Access.TTL) * time.Hour

	return &JWT{
		Access: &JWTAccessConfig{
			Secret:    dependencies.Config.Security.JWT.Access.Secret,
			ExpiresIn: expiresIn,
		},
	}
}

func (j *JWT) Create(payload JWTPayload) (string, time.Time) {
	stringUserId := strconv.Itoa(payload.UserID)
	expirationTime := time.Now().Add(j.Access.ExpiresIn)
	expiresAt := jwt.NewNumericDate(expirationTime)
	issuedAt := jwt.NewNumericDate(time.Now())

	claims := JWTClaims{
		UserID: stringUserId,
		Email:  payload.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.Access.Secret))

	if err != nil {
		panic(err)
	}

	return signedToken, expirationTime
}

func (j *JWT) Parse(token string) (*JWTPayload, error) {
	claims := JWTClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return []byte(j.Access.Secret), nil
	})

	if err != nil {
		return nil, errors.New(static.ErrorInvalidToken)
	}

	if !t.Valid {
		return nil, errors.New(static.ErrorInvalidToken)
	}

	intUserId, _ := strconv.Atoi(claims.UserID)

	payload := &JWTPayload{
		UserID: intUserId,
		Email:  claims.Email,
	}

	return payload, nil
}
