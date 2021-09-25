package helper

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	// ENCRYPTION_DEFAULT_KEY
	ENCRYPTION_DEFAULT_KEY = "6368616e676520746869732070617373776f726420746f206120736563726574"
	FIXED_IV               = "zLHVAb0SfJhZH/lR"
)

// SimpleEncryptLogs - Encrypt string using AES with block size of 16
func SimpleEncryptLogs(text string) string {

	if os.Getenv("ENCRYPT_LOGS_ENABLED") != "1" {
		return text
	}

	return SimpleEncryptString(text, "")
}

// SimpleEncryptString - Encrypt string using AES with block size of 16
func SimpleEncryptString(text string, iv string) string {

	var (
		key   []byte
		nonce []byte
	)

	if os.Getenv("ENCRYPTION_HELPER_KEY") != "" {
		key, _ = hex.DecodeString(os.Getenv("ENCRYPTION_HELPER_KEY"))
	} else {
		key, _ = hex.DecodeString(ENCRYPTION_DEFAULT_KEY)
	}

	plaintext := []byte(text)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	if iv == "" {
		nonce = make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}
	} else {
		nonce, _ = base64.StdEncoding.DecodeString(iv)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	t := aesgcm.Seal(nil, nonce, plaintext, nil)

	// pass the nonce for decryption. Its the last 12 bytes of the hex string
	return fmt.Sprintf("%x%x", t, nonce)
}

// SimpleDecryptString - Decrypt string using AES with block size of 16
func SimpleDecryptString(ciphertext string) string {

	var (
		key, cbytes, nonce []byte
		block              cipher.Block
		err                error
	)

	if os.Getenv("ENCRYPTION_HELPER_KEY") != "" {
		key, _ = hex.DecodeString(os.Getenv("ENCRYPTION_HELPER_KEY"))
	} else {
		key, _ = hex.DecodeString(ENCRYPTION_DEFAULT_KEY)
	}

	// trim last 24 character to get the cipher text
	cbytes, _ = hex.DecodeString(ciphertext[0 : len(ciphertext)-24])
	// get the last 24 character which is the nonce
	nonce, _ = hex.DecodeString(ciphertext[len(ciphertext)-24:])
	block, err = aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, cbytes, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func SimpleEncryptStringFixedIV(text string) string {
	return SimpleEncryptString(text, FIXED_IV)
}

//Encrypted data
func Encrypt(data, key []byte) ([]byte, error) {
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 16)
	content := PKCS5Padding(data, aesBlockEncrypter.BlockSize())
	encrypted := make([]byte, len(content))
	aesEncrypter := cipher.NewCBCEncrypter(aesBlockEncrypter, nonce)
	aesEncrypter.CryptBlocks(encrypted, content)
	return encrypted, nil
}

//Decrypt the data
func Decrypt(src, key []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err = aes.NewCipher(key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, []byte(FIXED_IV))
	aesDecrypter.CryptBlocks(decrypted, src)
	return PKCS5Trimming(decrypted), nil
}

/**
PKCS5 packaging
*/
func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

/*
 Unpack
*/
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

/*
 Returns encrypted string using sha256WithRSA
*/
func EncryptStringUsingsha256WithRSA(data []byte) (encodedSig string, err error) {

	var (
		keyPath string
	)

	if os.Getenv("ENCRYPTION_KEY_PATH") != "" {
		keyPath = os.Getenv("ENCRYPTION_KEY_PATH")
	} else {
		return "", errors.New("Key not found")
	}

	prvKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode([]byte(prvKey))
	if block == nil {
		return "", errors.New("Failed to parse root certificate PEM")
	}

	parseResult, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	key := parseResult.(*rsa.PrivateKey)

	ha := sha256.New()
	ha.Write([]byte(data))
	digest := ha.Sum(nil)

	s, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if err != nil {
		return "", errors.New("failed to sign:" + err.Error())
	}

	encodedSig = base64.StdEncoding.EncodeToString(s)

	return encodedSig, nil
}
