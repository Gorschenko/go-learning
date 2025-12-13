package jwt_test

import (
	"test/packages/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	const email = "test@test.ru"
	jwtService := jwt.NewJWT("your-256-bit-secret-key-here-123456")

	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})

	if err != nil {
		t.Fatal(err)
	}

	isValid, data := jwtService.Parse(token)

	if !isValid {
		t.Fatal("Token is invalid")
	}

	if data.Email != email {
		t.Fatalf("Email %s email not equal %s", email, data.Email)
	}
}
