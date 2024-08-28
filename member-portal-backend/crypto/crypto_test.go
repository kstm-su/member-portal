package crypto

import (
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	password := "password"
	encryptedPassword := passwordEncryptWithParams(password, "salt-salt", "pepper", 64*1024, 1, 4, 32)
	if encryptedPassword != "$argon2id$v=19$m=65536,t=4,p=1$c2FsdC1zYWx0$6JNmlGvpjNKYpQNSJdGNfAJQ7+upIXwebdDMWcJf30g" {
		t.Errorf("PasswordEncrypt failed, expected $argon2id$v=19$m=65536,t=4,p=1$c2FsdC1zYWx0$6JNmlGvpjNKYpQNSJdGNfAJQ7+upIXwebdDMWcJf30g, got %s", encryptedPassword)
	}
}

func TestVerifyPassword(t *testing.T) {
	password := "password"
	encryptedPassword := "$argon2id$v=19$m=65536,t=4,p=1$cmFmamRhc2poZmRzYWpoa2xhc2Rma2po$i39rjDYap4n6eA2XhCusp5wHGPKBBpM0Lg9S82kw0Ec"
	if !VerifyPassword(encryptedPassword, password) {
		t.Errorf("VerifyPassword failed, expected true, got false")
	}

	wrongPassword := "wrongpassword"
	if VerifyPassword(encryptedPassword, wrongPassword) {
		t.Errorf("VerifyPassword failed, expected false, got true")
	}
}

func TestGenerateRandomString(t *testing.T) {
	length := 30
	randomString := GenerateRandomString(length)
	if len(randomString) != length {
		t.Errorf("GenerateRandomString failed, expected length %d, got %d", length, len(randomString))
	}
}
