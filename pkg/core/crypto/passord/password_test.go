package passord_test

import (
	"testing"

	"github.com/chengyu914001/go-common/pkg/core/crypto/passord"
)

func TestVerifyUserPassword(t *testing.T) {
	exceptedPassword := "test password!"
	hashedPassword, err := passord.HashUserPassword(exceptedPassword)
	if err != nil {
		t.Error(err)
	}

	verified, err := passord.VerifyUserPassword(exceptedPassword, hashedPassword)
	if err != nil {
		t.Error(err)
	}

	if !verified {
		t.Error("verify failed")
	}
}
