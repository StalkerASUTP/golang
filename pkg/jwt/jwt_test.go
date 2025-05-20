package jwt_test

import (
	"go/adv-api/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("78cd9185a4d4234d1a9c5e3f6c3ea21590e8c0282ab34ed51990bfc3e72e4906")
	token, err := jwtService.Create(jwt.JWTData{
		Email: "a@a.ru",
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
