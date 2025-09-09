package token

import (
	"time"

	"github.com/chengyu914001/go-common/pkg/core/permission"
	"github.com/chengyu914001/go-common/pkg/core/sysconst"
	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenTTL = 10 * time.Minute
)

type accessTokenClaims struct {
	Scopes permission.Permission `json:"scp"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	userID string,
	scopes permission.Permission,
	audiences []string,
) (string, error) {
	now := time.Now()
	exp := now.Add(accessTokenTTL)
	registeredClaims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(now),
		Issuer:    sysconst.GetIssuerName(),
	}
	if len(audiences) > 0 {
		registeredClaims.Audience = audiences
	}
	claims := accessTokenClaims{
		Scopes:           scopes,
		RegisteredClaims: registeredClaims,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signed, err := token.SignedString(privateKey)
	if err != nil {
		return "", ErrGenerateToken
	}
	return signed, nil
}

func ValidateAccessToken(tokenStr string) (*accessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &accessTokenClaims{}, getPublicKey)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(*accessTokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
