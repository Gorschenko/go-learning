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
		AccessConfig: &JWTAccessConfig{
			Secret:    dependencies.Config.Security.JWT.Access.Secret,
			ExpiresIn: expiresIn,
		},
	}
}

func (j *JWT) Create(data JWTDataToCreate) *JWTToken {
	stringUserId := strconv.Itoa(data.UserID)
	expirationTime := time.Now().Add(j.AccessConfig.ExpiresIn)
	expiresAt := jwt.NewNumericDate(expirationTime)
	issuedAt := jwt.NewNumericDate(time.Now())

	claims := JWTClaims{
		UserID: stringUserId,
		Email:  data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.AccessConfig.Secret))

	if err != nil {
		panic(err)
	}

	return &JWTToken{
		Token:          signedToken,
		ExpirationTime: expirationTime,
	}
}

func (j *JWT) Parse(token string) (*JWTDataToCreate, error) {
	claims := JWTClaims{}
	t, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return []byte(j.AccessConfig.Secret), nil
	})

	if err != nil {
		return nil, errors.New(static.ErrorInvalidToken)
	}

	if !t.Valid {
		return nil, errors.New(static.ErrorInvalidToken)
	}

	intUserId, _ := strconv.Atoi(claims.UserID)

	payload := &JWTDataToCreate{
		UserID: intUserId,
		Email:  claims.Email,
	}

	return payload, nil
}
