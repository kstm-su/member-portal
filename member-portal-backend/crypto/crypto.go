package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/kstm-su/Member-Portal/backend/config"
	"golang.org/x/crypto/argon2"
	_ "golang.org/x/crypto/argon2"
	"math/big"
	"strings"
)

func PasswordEncrypt(password string) string {
	salt := []byte(GenerateRandomString(30))
	pepper := config.Cfg.Password.Pepper

	memory := 64 * 1024
	iterations := 4
	parallelism := 1
	keyLen := 32

	encoded := passwordEncryptWithParams(password, string(salt), pepper, memory, parallelism, iterations, keyLen)

	return encoded
}

func passwordEncryptWithParams(password string, salt string, pepper string, memory int, parallelism int, iterations int, keyLen int) string {
	withPepper := password + pepper

	hash := argon2.IDKey([]byte(withPepper), []byte(salt), uint32(iterations), uint32(memory), uint8(parallelism), uint32(keyLen))

	b64Salt := base64.RawStdEncoding.EncodeToString([]byte(salt))
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)

	return encodedHash
}

func VerifyPassword(hash string, password string) bool {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	withPepper := password + config.Cfg.Password.Pepper

	genHash := argon2.IDKey([]byte(withPepper), salt, uint32(4), uint32(64*1024), uint8(1), uint32(32))

	if bytes.Equal(decodedHash, genHash) {
		return true
	}

	return false
}

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return ""
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}
