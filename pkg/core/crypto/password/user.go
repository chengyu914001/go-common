package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

func HashUserPassword(password string) (string, error) {
	lock.Lock()
	defer lock.Unlock()

	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	key := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)

	// len: 22 + 43
	res := base64.RawStdEncoding.EncodeToString(salt) + base64.RawStdEncoding.EncodeToString(key)
	return res, nil
}

func VerifyUserPassword(password string, hashedPassword string) (bool, error) {
	lock.Lock()
	defer lock.Unlock()

	salt, err := base64.RawStdEncoding.DecodeString(hashedPassword[:22])
	if err != nil {
		return false, err
	}

	correctKey, err := base64.RawStdEncoding.DecodeString(hashedPassword[22:])
	if err != nil {
		return false, err
	}

	queryKey := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)
	return subtle.ConstantTimeCompare(queryKey, correctKey) == 1, nil
}
