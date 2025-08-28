package crypto

import (
	"context"

	"github.com/chengyu914001/go-common/pkg/core/cpu"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(ctx context.Context, password string) (string, error) {
	res := []byte{}
	err := cpu.DoLimitRun(ctx, func() error {
		var innerErr error
		res, innerErr = bcrypt.GenerateFromPassword([]byte(password), 12)
		return innerErr
	})
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func CheckPasswordHash(ctx context.Context, password, hash string) (bool, error) {
	res := false
	err := cpu.DoLimitRun(context.Background(), func() error {
		innerErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		res = innerErr == nil
		return nil
	})
	return res, err
}
