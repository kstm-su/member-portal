package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/kstm-su/Member-Portal/backend/config"
	"log"
	"os"
)

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenKey(bits int, config config.Config) error {
	baseDir := config.File.Base
	keyDir := baseDir + "/key"

	// キー保存用のディレクトリを作成
	err := os.MkdirAll(keyDir, 0700)
	if err != nil {
		return err
	}
	privateKeyFileName := keyDir + "/private_key.pem"
	pubKeyFileName := keyDir + "/public_key.pem"

	// 秘密鍵および公開鍵が既に存在する場合は終了
	if _, err := os.Stat(privateKeyFileName); err == nil {
		log.Println("秘密鍵が既に存在します")
		return nil
	}

	if _, err := os.Stat(pubKeyFileName); err == nil {
		return nil
	}

	log.Println("秘密鍵が存在しないため新規作成します")
	// 秘密鍵を生成
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	// 秘密鍵をファイルに保存
	privateKeyFile, err := os.Create(keyDir + "/private_key.pem")
	if err != nil {
		return err
	}
	defer func(privateKeyFile *os.File) {
		err := privateKeyFile.Close()
		if err != nil {
			panic(err)
		}
	}(privateKeyFile)

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// 公開鍵をファイルを取得
	publicKey := &privateKey.PublicKey

	// 公開鍵をファイルに保存
	publicKeyFile, err := os.Create(keyDir + "/public_key.pem")
	if err != nil {
		return err
	}
	defer func(publicKeyFile *os.File) {
		err := publicKeyFile.Close()
		if err != nil {
			panic(err)
		}
	}(publicKeyFile)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		return err
	}

	return nil
}

func GetKeys(config config.Config) Key {
	var keys Key
	// 秘密鍵を読み込む
	privateKeyFile, err := os.ReadFile(config.File.Base + "/key/private_key.pem")
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(privateKeyFile)
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	keys.PrivateKey = privateKey

	// 公開鍵を読み込む
	publicKeyFile, err := os.ReadFile(config.File.Base + "/key/public_key.pem")
	if err != nil {
		panic(err)
	}
	block, _ = pem.Decode(publicKeyFile)
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	keys.PublicKey = publicKey.(*rsa.PublicKey)

	return keys
}
