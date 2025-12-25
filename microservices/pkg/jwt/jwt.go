package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(dependencies JWTDependencies) *JWT {
	return &JWT{
		Access: &JWTToken{
			Secret:   dependencies.Config.Security.JWT.Access.Secret,
			TTLHours: dependencies.Config.Security.JWT.Access.TTLHours,
		},
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": data.UserID,
		"email":  data.Email,
	})

	signedToken, err := token.SignedString([]byte(j.Access.Secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWT) Parse(token string) (*JWTData, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(j.Access.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	userID := t.Claims.(jwt.MapClaims)["userID"].(int)
	email := t.Claims.(jwt.MapClaims)["email"].(string)

	data := &JWTData{
		UserID: userID,
		Email:  email,
	}

	return data, nil
}
