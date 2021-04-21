package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, _ := newCipherBlock(key)
	return cipher.NewCFBEncrypter(block, iv), nil
}

// EncryptWriter return a writer that write encrypted data to the original writer.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, _ := encryptStream(key, iv)
	n, err := w.Write(iv)
	err = unabletowrite(n, iv, err)
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func unabletowrite(n int, iv []byte, err error) error {
	if len(iv) != n || err != nil {
		return errors.New("Unable to write IV into writer")
	}
	return nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, _ := newCipherBlock(key)
	return cipher.NewCFBDecrypter(block, iv), nil
}

// DecryptReader will return a reader that will decrypt data from the provided reader 
//and give the user a way to read that data as it if was not encrypted.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	err = unabletoread(n,iv,err)
	stream, _ := decryptStream(key, iv)
	return &cipher.StreamReader{S: stream, R: r}, nil
}
func unabletoread(n int, iv []byte, err error) error {
	if n < len(iv) || err != nil {
		return  errors.New("encrypt: unable to read the full iv")
	}
	return nil
}

func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
