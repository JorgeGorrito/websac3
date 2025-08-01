package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	Sub         uint
	Username    string
	Role        string
	Permissions map[string][]string
}

type RefreshTokenClaims struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}

type TokenJWTResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Generator interface {
	GenerateAccessToken(tr AccessTokenClaims) (string, error)
	ValidateRefreshToken(tokenStr string) (*RefreshTokenClaims, error)
	GenerateTokens(tr AccessTokenClaims, rr RefreshTokenClaims) (TokenJWTResponse, error)
}

type generator struct {
	accessTokenSecret    string
	refrestTokenSecret   string
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewGenerator(accessTokenSecret, refreshTokenSecret string, accessTokenDuration, refreshTokenDuration time.Duration) Generator {
	return &generator{
		accessTokenSecret:    accessTokenSecret,
		refrestTokenSecret:   refreshTokenSecret,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (g *generator) GenerateTokens(tr AccessTokenClaims, rr RefreshTokenClaims) (TokenJWTResponse, error) {
	accessToken, err := g.GenerateAccessToken(tr)
	if err != nil {
		return TokenJWTResponse{}, err
	}

	refreshToken, err := g.GenerateRefreshToken(rr)
	if err != nil {
		return TokenJWTResponse{}, err
	}

	return TokenJWTResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (g *generator) GenerateAccessToken(tr AccessTokenClaims) (string, error) {
	claims := jwt.MapClaims{
		"sub":         tr.Sub,
		"username":    tr.Username,
		"role":        tr.Role,
		"permissions": tr.Permissions,
		"exp":         time.Now().Add(g.accessTokenDuration).Unix(),
		"iat":         time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := g.accessTokenSecret
	if secret == "" {
		return "", jwt.ErrSignatureInvalid
	}

	return token.SignedString([]byte(secret))
}

func (g *generator) GenerateRefreshToken(tr RefreshTokenClaims) (string, error) {
	claims := jwt.MapClaims{
		"sub": tr.Sub,
		"exp": time.Now().Add(g.refreshTokenDuration).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := g.refrestTokenSecret
	if secret == "" {
		return "", jwt.ErrSignatureInvalid
	}

	return token.SignedString([]byte(secret))
}

func (g *generator) ValidateRefreshToken(tokenStr string) (*RefreshTokenClaims, error) {
	claims := &RefreshTokenClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(g.refrestTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	return claims, nil
}
