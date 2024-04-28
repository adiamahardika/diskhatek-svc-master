package models

import "github.com/golang-jwt/jwt"

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	SignatureKey string `json:"signature_key"`
	Username     string `json:"username"`
	BranchKey    string `json:"branchKey"`
	jwt.StandardClaims
}
