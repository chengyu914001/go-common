package token

import (
	"time"

	"crypto/rand"
	"math/big"

	"github.com/chengyu914001/go-common/pkg/core/id"
	"github.com/chengyu914001/go-common/pkg/core/sysconst"
	"github.com/golang-jwt/jwt/v5"
)

const accessTokenTTL = 10 * time.Minute

var randAccessTokenDurationGenerator *big.Int

func init() {
	var err error
	randAccessTokenDurationGenerator, err = rand.Int(rand.Reader, big.NewInt(int64(time.Second+1)))
	if err != nil {
		panic(err)
	}
}

func generateAccessTokenRandDuration() time.Duration {
	return time.Duration(randAccessTokenDurationGenerator.Int64())
}

type accessTokenClaims struct {
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string, audiences []string) (string, error) {
	now := time.Now().UTC()
	exp := now.Add(accessTokenTTL + generateAccessTokenRandDuration())
	claims := accessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    sysconst.GetIssuerName(),
			Subject:   userID,
			Audience:  jwt.ClaimStrings(audiences),
			ExpiresAt: jwt.NewNumericDate(exp),
			ID:        id.GenerateTokenID(),
		},
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
