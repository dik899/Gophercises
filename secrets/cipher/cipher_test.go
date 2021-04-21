package cipher

import (
	
	"os"
	"path/filepath"
	"testing"
	"bytes"
	"crypto/aes"
	"errors"
   homedir "github.com/mitchellh/go-homedir"
)

func TestEncryptWriter(t *testing.T) {
	var b bytes.Buffer
	key := "abc"
	_, err := EncryptWriter(key, &b)
	if err != nil {
		t.Errorf("Error %v", err)
	}
}
func TestDecryptReader(t *testing.T) {
	home, _ := homedir.Dir()
	fp := filepath.Join(home, "secret.txt")

	f, _ := os.Open(fp)
	defer f.Close()
	
	_, err := DecryptReader("abc", f)
	if err != nil {
		t.Errorf("Error : %v ", err)
	}
}
func TestUnabletoWrite(t *testing.T) {
	iv := make([]byte, aes.BlockSize)
	err := unabletowrite(10, iv, errors.New("test"))
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
func TestUnabletoRead(t *testing.T) {
	iv := make([]byte, aes.BlockSize)
	err := unabletoread(10, iv, errors.New("test"))
	if err == nil {
		t.Error("Expected error but got no error")
	}
}
