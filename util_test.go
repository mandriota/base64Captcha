package base64Captcha

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_randomBytes(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomBytes(i)
		if len(digti) != i {
			t.Error("failed")
		}
	}
}

func Test_randomBytesMod(t *testing.T) {
	for i := 1; i < 10; i++ {
		digti := randomBytesMod(i, 'c')
		if len(digti) != i {
			t.Error("failed")
		}
	}
}

func Test_itemWriteFile(t *testing.T) {
	//todo:::
}

func Test_pathExists(t *testing.T) {
	p := filepath.Join(os.TempDir(), RandomId())
	if pathExists(p) {
		t.Error("failed")
	}
	_ = os.MkdirAll(p, os.ModePerm)
	defer os.Remove(p)

	if !pathExists(p) {
		t.Error("failed")
	}
}
