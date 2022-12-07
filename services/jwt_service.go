package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	issure string
}

func NewJWTService() *jwtService {
	return &jwtService{
		issure: "user-api",
	}
}

type Claim struct {
	Sum int `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id int) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.issure))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}
		return []byte(s.issure), nil
	})

	if err != nil {
		fmt.Println("Error on ValidateToken ", err)
	}

	return err == nil
}
