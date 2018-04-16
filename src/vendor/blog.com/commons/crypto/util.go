package ymtcrypto

import (
	"crypto/aes"
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

func RandomBytes(len int) []byte {
	s := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, len)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result[i] = s[r.Intn(62)]
	}

	return result
}

/**
 * 1.generate random 32 byte salt S
 * 2.encrypt S with rsa to R1
 * 3.use left aes.BlockSize byte of S as the aes key K
 * 4.encrypt plain text with aes to R2
 * 5.concat base64-encoded R1 and base64-encoded R2 with =
 */
func YMTPayEncrypt(r *RSAManager, plain_text []byte) (string, error) {
	aes_key := RandomBytes(32)
	aes_key_encrypt, err := r.Encrypt(aes_key)
	if err != nil {
		return "", err
	}

	aes_key_16 := aes_key[:aes.BlockSize]
	a, _ := NewAESManager(aes_key_16)
	ciphertext, err := a.Encrypt(plain_text)
	if err != nil {
		return "", err
	}

	aes_key_encrypt_b64 := base64.StdEncoding.EncodeToString(aes_key_encrypt)
	ciphertext_b64 := base64.StdEncoding.EncodeToString(ciphertext)
	return aes_key_encrypt_b64 + "=" + ciphertext_b64, nil
}

func YMTPayDecrypt(r *RSAManager, ciphertext string) ([]byte, error) {
	tuple := strings.SplitAfterN(ciphertext, "=", 2)
	rsa_ciphertext, err := base64.StdEncoding.DecodeString(tuple[0])
	if err != nil {
		return nil, err
	}

	aes_ciphertext, err := base64.StdEncoding.DecodeString(strings.TrimLeft(tuple[1], "="))
	if err != nil {
		return nil, err
	}

	aes_key, err := r.Decrypt(rsa_ciphertext)
	if err != nil {
		return nil, err
	}
	aes_key_16 := aes_key[:16]
	a, _ := NewAESManager(aes_key_16)

	plain_txt, err := a.Decrypt(aes_ciphertext)
	return plain_txt, err
}
