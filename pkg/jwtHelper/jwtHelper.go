package jwtHelper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

type JWT struct {
	priKey []byte
	pubKey []byte
}

func NewJWT() *JWT {
	prvKey, err := os.ReadFile("keys/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := os.ReadFile("keys/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	return &JWT{priKey: prvKey, pubKey: pubKey}
}

func (j JWT) Create(ttl time.Duration, username string, userId string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.priKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"issuer":   "Alfa gov",
		"exp":      time.Now().Add(ttl).Unix(),
	}).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (j JWT) Validate(token string) (string, string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", "", fmt.Errorf("validate: parse key: %w", err)
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return "", "", fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", fmt.Errorf("validate: invalid")
	}

	user, ok := claims["username"].(string)
	if !ok || user == "" {
		return "", "", fmt.Errorf("validate: invalid")
	}

	userId, ok := claims["userId"].(string)
	if !ok || userId == "" {
		return "", "", fmt.Errorf("validate: invalid")
	}

	return user, userId, nil
}
