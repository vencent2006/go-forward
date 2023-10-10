package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthClaim struct {
	UserId uint64 `json:"userId"`
	jwt.StandardClaims
}

var secret = []byte("It is my secret")

const issuer = "TEST001"

const TokenExpireDuration = 2 * time.Hour

func GenToken(userId uint64) (string, error) {
	c := AuthClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func findSecret(tk *jwt.Token) (interface{}, error) {
	claim, ok := tk.Claims.(*AuthClaim)
	if !ok {
		return nil, errors.New("Invalid claim ")
	}
	if claim.Issuer == issuer {
		return secret, nil
	}
	return nil, fmt.Errorf("invalid issuer:%s", claim.Issuer)
}

func ParseToken(tokenStr string) (*AuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaim{}, findSecret)
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*AuthClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("Invalid token ")
}
