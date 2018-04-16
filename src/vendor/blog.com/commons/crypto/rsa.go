package ymtcrypto

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"

	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

type RSAManager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func loadFile(p string) ([]byte, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func loadPrivateKey(pri_key_path string) (*rsa.PrivateKey, error) {
	b, err := loadFile(pri_key_path)
	if err != nil {
		return nil, errors.New("load private key file failed")
	}

	pri_block, _ := pem.Decode(b)
	if pri_block == nil {
		return nil, errors.New("load private key pem failed")
	}

	pri_key, err := x509.ParsePKCS1PrivateKey(pri_block.Bytes)
	if err != nil {
		return nil, errors.New("load private key failed")
	}
	//prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
	//if err != nil {
	//	return nil, err
	//}
	//priKey = prkI.(*rsa.PrivateKey)

	return pri_key, nil
}

func loadPublicKey(pub_key_path string) (*rsa.PublicKey, error) {
	b, err := loadFile(pub_key_path)
	if err != nil {
		return nil, errors.New("load public key file failed")
	}

	pub_block, _ := pem.Decode(b)
	if pub_block == nil {
		return nil, errors.New("load public key pem failed")
	}

	pub_key, err := x509.ParsePKIXPublicKey(pub_block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub_key.(*rsa.PublicKey), nil
}

func NewRSAManager(pri_key_path, pub_key_path string) (*RSAManager, error) {
	pri_key, err := loadPrivateKey(pri_key_path)
	if err != nil {
		return nil, errors.New("load private key failed")
	}
	pub_key, err := loadPublicKey(pub_key_path)
	if err != nil {
		return nil, errors.New("load pub key failed")
	}

	return &RSAManager{
		privateKey: pri_key,
		publicKey:  pub_key,
	}, nil
}

func (this *RSAManager) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, plaintext)
}
func (this *RSAManager) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, ciphertext)
}

func (this *RSAManager) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
}

func (this *RSAManager) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
}
