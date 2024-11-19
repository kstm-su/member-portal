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

func PasswordEncrypt(password string, config *config.Config) string {
	salt := []byte(GenerateRandomString(30))

	pepper := config.Password.Pepper

	memory := 64 * 1024
	iterations := 4
	threads := 1
	keyLen := 32

	encoded := passwordEncryptWithParams(password, string(salt), pepper, memory, threads, iterations, keyLen)

	return encoded
}

func passwordEncryptWithParams(password string, salt string, pepper string, memory int, threads int, iterations int, keyLen int) string {
	// パスワードとpepperを連結
	withPepper := password + pepper

	// ハッシュ値を生成
	hash := argon2.IDKey([]byte(withPepper), []byte(salt), uint32(iterations), uint32(memory), uint8(threads), uint32(keyLen))

	// ハッシュ値とsaltをbase64エンコード
	b64Salt := base64.RawStdEncoding.EncodeToString([]byte(salt))
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// データベースに保存するための文字列を生成
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, iterations, threads, b64Salt, b64Hash)

	return encodedHash
}

func VerifyPassword(hash string, password string, config config.Config) bool {
	// 例: $argon2id$v=19$m=65536,t=4,p=1$c2FsdC1zYWx0$6JNmlGvpjNKYpQNSJdGNfAJQ7+upIXwebdDMWcJf30g
	// このような形式の文字列をパースする
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false
	}

	// パースした文字列から必要な情報を取り出す
	// saltがbase64エンコードされているのでデコードする
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	// ハッシュ値もbase64エンコードされているのでデコードする
	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// ハッシュ値を生成する際に使用したパラメータを取り出す
	// m=65536,t=4,p=1
	var iterator int
	var memory uint32
	var threads uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterator, &threads)

	if err != nil {
		panic(err.Error())
	}
	// ハッシュ値の長さを取得
	keyLen := len(decodedHash)

	// パスワードとpepperを連結してハッシュ値を生成
	withPepper := password + config.Password.Pepper

	// ハッシュ値を生成
	genHash := argon2.IDKey([]byte(withPepper), salt, uint32(iterator), memory, threads, uint32(keyLen))

	// 生成したハッシュ値とデータベースに保存されているハッシュ値を比較
	return bytes.Equal(decodedHash, genHash)
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
