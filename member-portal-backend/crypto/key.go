package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"github.com/google/uuid"
	"github.com/kstm-su/Member-Portal/backend/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"log/slog"
	"math/big"
	"os"
	"time"
)

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func Init(config config.Config) {
	slog.Info("キーペアの生成を開始します")
	err := genKey(config)
	if err != nil {
		return
	}
	slog.Info("キーペアの生成が完了しました")

	slog.Info("キーペアより証明書を取得します")
	keys := GetKeys(config)
	err = generateCertificate(keys.PrivateKey, keys.PublicKey)
	if err != nil {
		panic("証明書の取得に失敗しました")
	}
	slog.Info("証明書の取得が完了しました")

	slog.Info("jwks.jsonを生成します")
	generateJWKs(config)
	slog.Info("jwks.jsonの生成が完了しました")
}

func genKey(config config.Config) error {
	baseDir := config.File.Base
	keyDir := baseDir + "/key"

	// キー保存用のディレクトリを作成
	err := os.MkdirAll(keyDir, 0700)
	if err != nil {
		return err
	}
	privKeyFileName := keyDir + "/private_key.pem"
	pubKeyFileName := keyDir + "/public_key.pem"

	// 秘密鍵および公開鍵が既に存在する場合は終了
	if _, err := os.Stat(privKeyFileName); err == nil {
		if _, err := os.Stat(pubKeyFileName); err == nil {
			slog.Info("キーペアが既に存在します")
			return nil
		}
	}

	slog.Info("キーペアが存在しないため新規作成します")
	// 秘密鍵を生成
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
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

func generateCertificate(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) error {
	startDate := time.Now()
	endDate := startDate.AddDate(1, 0, 0) // 1 year validity
	serialNumber := big.NewInt(time.Now().Unix())
	subject := pkix.Name{
		CommonName: "Test Certificate",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		Issuer:       subject, // self-signed
		NotBefore:    startDate,
		NotAfter:     endDate,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey, privateKey)
	if err != nil {
		return err
	}

	certFile, err := os.Create("certificate.pem")
	if err != nil {
		return err
	}
	defer func(certFile *os.File) {
		err := certFile.Close()
		if err != nil {
			panic("証明書の書き込みに失敗しました")
		}
	}(certFile)

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	if err != nil {
		return err
	}

	return nil
}

func generateJWKs(config config.Config) {
	jwksFile := config.File.Base + "/key/jwks.json"

	//既にファイルが存在している場合
	if _, err := os.Stat(jwksFile); err == nil {
		slog.Info("jwks.jsonが既に存在します")
		return
	}

	//そうでない場合
	file, err := os.Create(jwksFile)
	if err != nil {
		slog.Warn("jwks.jsonの生成に失敗しました")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("jwks.jsonの書き込みに失敗しました")
		}
	}(file)

	privateKey := GetKeys(config).PrivateKey

	key, err := jwk.PublicKeyOf(privateKey)
	if err != nil {
		slog.Warn("jwksのpublickey生成に失敗しました")
	}

	uniqueId := uuid.New()

	_ = key.Set(jwk.KeyIDKey, uniqueId)
	_ = key.Set(jwk.KeyUsageKey, jwk.ForSignature)
	_ = key.Set(jwk.KeyTypeKey, jwa.RS256)

	// Create a JWK set and add the JWK to the set
	jwkSet := jwk.NewSet()
	err = jwkSet.AddKey(key)
	if err != nil {
		return
	}

	encoded, _ := json.Marshal(jwkSet)

	// Save the JWK set to a file
	_, err = file.Write(encoded)
	if err != nil {
		slog.Warn("jwks.jsonの生成に失敗しました")
		return
	}
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
		panic("秘密鍵の読み込みに失敗しました")
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
		panic("公開鍵の読み取りに失敗しました")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	keys.PublicKey = publicKey.(*rsa.PublicKey)

	return keys
}
