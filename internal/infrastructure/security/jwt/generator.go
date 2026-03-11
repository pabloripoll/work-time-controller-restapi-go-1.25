package jwt

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Generator struct {
	secretKey []byte
	issuer    string
}

func NewGenerator(secretKey, issuer string) *Generator {
	return &Generator{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}

func (g *Generator) GenerateToken(userID int64, email, role string, duration time.Duration) (string, time.Time, error) {
	expiresAt := time.Now().Add(duration)
	userIDStr := strconv.FormatInt(userID, 10)

	claims := Claims{
		UserID: userIDStr,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    g.issuer,
			Subject:   userIDStr,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.secretKey)

	return tokenString, expiresAt, err
}

func (g *Generator) GenerateAccessToken(userID int64, email, role string) (string, time.Time, error) {
	return g.GenerateToken(userID, email, role, 15*time.Minute) // 15 minutes
}

func (g *Generator) GenerateRefreshToken(userID int64, email, role string) (string, time.Time, error) {
	return g.GenerateToken(userID, email, role, 7*24*time.Hour) // 7 days
}
