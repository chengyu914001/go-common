package password_test

import (
	"testing"

	"github.com/chengyu914001/go-common/pkg/core/crypto/password"
)

func TestVerifyUserPassword(t *testing.T) {
	exceptedPassword := "test password!"
	hashedPassword, err := password.HashUserPassword(exceptedPassword)
	if err != nil {
		t.Error(err)
	}

	verified, err := password.VerifyUserPassword(exceptedPassword, hashedPassword)
	if err != nil {
		t.Error(err)
	}

	if !verified {
		t.Error("verify failed")
	}
}
