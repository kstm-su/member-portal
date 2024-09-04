package crypto

import (
	"github.com/kstm-su/Member-Portal/backend/config"
	"os"
	"testing"
)

func TestGenKey(t *testing.T) {
	var cfg = config.Config{}
	cfg.File.Base = "/tmp"
	err := GenKey(2048, cfg)
	if err != nil {
		t.Errorf("GenKey failed, expected nil, got %v", err)
	}

	// 生成されたキーが存在するか確認
	_, err = os.Stat("/tmp/key/private_key.pem")
	if err != nil {
		t.Errorf("GenKey failed, private_key.pem not found")
	}

	_, err = os.Stat("/tmp/key/public_key.pem")
	if err != nil {
		t.Errorf("GenKey failed, public_key.pem not found")
	}
}
