package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Интерфейс для использования и моков
type JWTManager interface {
	GenerateAccessToken(userID int64, role string) (string, error)
	GenerateRefreshToken(userID int64, role string) (string, error)
	Verify(tokenStr string) (*Claims, error)
	RefreshTTL() time.Duration
}

type jwtManager struct {
	secretKey  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// NewJWTManager возвращает интерфейс JWTManager
func NewJWTManager(secretKey string, accessTTL, refreshTTL time.Duration) JWTManager {
	return &jwtManager{
		secretKey:  secretKey,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

func (jm *jwtManager) GenerateAccessToken(userID int64, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jm.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *jwtManager) GenerateRefreshToken(userID int64, role string) (string, error) {
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jm.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.secretKey))
}

func (jm *jwtManager) Verify(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jm.secretKey), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func (jm *jwtManager) RefreshTTL() time.Duration {
	return jm.refreshTTL
}
