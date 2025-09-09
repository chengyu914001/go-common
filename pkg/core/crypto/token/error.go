package token

import "errors"

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrGenerateToken = errors.New("generate token error")
)
