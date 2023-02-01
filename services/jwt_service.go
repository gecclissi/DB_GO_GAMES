package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
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
	IdJogador int `json:"id_jogador"`
	jwt.StandardClaims
}

func (s *jwtService) GetClaimFromToken(toke string) (Claim, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(toke, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(s.issure), nil
	})
	if err != nil {
		return Claim{}, err
	}
	claims := token.Claims.(*Claim)

	return *claims, nil
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
