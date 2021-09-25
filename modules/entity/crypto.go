package entity

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
)

const (
	iVDelimeter = "|"
)

var (
	privateKeyValue []byte
	publicKeyValue  []byte
)

// init function ...
func init() {
	privateKeyValue, _ = ioutil.ReadFile(os.Getenv(constant.RsaPrivateKeyLocation))
	publicKeyValue, _ = ioutil.ReadFile(os.Getenv(constant.RsaPublicKeyLocation))
}

// AES_Encrypt function ...
func AES_Encrypt(plainText string, key []byte) (cipherText []byte, err error) {
	return AES_256_CBC_Encrypt(plainText, key)
}

// AES_Decrypt function ...
func AES_Decrypt(cipherText []byte, key []byte) (plainText string, err error) {
	return AES_256_CBC_Decrypt(cipherText, key)
}

// AES_256_GCM_Encrypt function ...
func AES_256_GCM_Encrypt(plainText string, key []byte) (cipherText []byte, err error) {
	// generate a new aes cipher using our key
	c, err := aes.NewCipher(key)

	// if there are any errors, handle them
	if err != nil {
		return nil, err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)

	// if any error generating new GCM, handle them
	if err != nil {
		return nil, err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	cipherText = gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return
}

// AES_256_GCM_Decrypt function ...
func AES_256_GCM_Decrypt(cipherText []byte, key []byte) (plainText string, err error) {
	// generate a new aes cipher using our key
	c, err := aes.NewCipher(key)

	// if there are any errors, handle them
	if err != nil {
		return "", err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)

	// if any error generating new GCM, handle them
	if err != nil {
		return "", err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonceSize := gcm.NonceSize()

	if len(cipherText) < nonceSize {
		return "", err
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	pt, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		return "", err
	}

	plainText = string(pt)

	return
}

// AES_256_GCM_Encrypt function ...
func AES_256_CBC_Encrypt(plainText string, b64Key []byte) (cipherText []byte, err error) {
	key, err := base64.RawStdEncoding.DecodeString(string(b64Key))

	if err != nil {
		return nil, err
	}

	plaintextBytes := []byte(plainText)

	// Adds padding `0x00` hex byte upto block size divisible.
	for len(plaintextBytes)%aes.BlockSize != 0 {
		plaintextBytes = append(plaintextBytes, byte(0x00))
	}

	// CBC mode works on blocks so plaintextBytess may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintextBytes is already of the correct length.
	if len(plaintextBytes)%aes.BlockSize != 0 {
		return nil, errors.New("plaintextBytes is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintextBytes)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	cipherText = []byte(fmt.Sprintf("%x", iv) + iVDelimeter + ConvertUnsafeBytesToBase64String(ciphertext[aes.BlockSize:]))
	cipherTextWithIV := []byte(ConvertUnsafeBytesToBase64String(cipherText))

	return cipherTextWithIV, err
}

// AES_256_CBC_Decrypt function ...
func AES_256_CBC_Decrypt(b64CipherText []byte, b64Key []byte) (plainText string, err error) {
	key, err := base64.RawStdEncoding.DecodeString(string(b64Key))

	if err != nil {
		return "", err
	}

	cipherTextWithIV, _ := ConvertBase64StringToUnsafeBytes(string(b64CipherText))

	if err != nil {
		return "", err
	}

	ctIvArray := strings.Split(string(cipherTextWithIV), iVDelimeter)
	iv, err := hex.DecodeString(ctIvArray[0])

	if err != nil {
		return "", err
	}

	cipherText := ctIvArray[1]
	cipherTextBytes, err := ConvertBase64StringToUnsafeBytes(cipherText)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	if len(cipherTextBytes) < aes.BlockSize {
		return "", errors.New("cipherTextBytes too short")
	}

	// CBC mode always works in whole blocks.
	if len(cipherTextBytes)%aes.BlockSize != 0 {
		return "", errors.New("cipherTextBytes is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(cipherTextBytes, cipherTextBytes)

	// Remove padding
	cipherTextBytes = bytes.TrimRight(cipherTextBytes, "\x00")

	// Remove newline
	decryptedText := strings.Replace(string(cipherTextBytes), "\n", "", -1)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.
	return decryptedText, err
}

// RSA_Encrypt function ...
func RSA_Encrypt(plainText []byte) ([]byte, error) {
	if len(publicKeyValue) == 0 {
		return nil, errors.New("Failed to read Public key")
	}

	block, _ := pem.Decode(publicKeyValue)

	if block == nil {
		return nil, errors.New("Public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

// RSA_Decrypt function ...
func RSA_Decrypt(ciphertext []byte) ([]byte, error) {
	if len(privateKeyValue) == 0 {
		return nil, errors.New("Failed to read Private key")
	}

	block, _ := pem.Decode(privateKeyValue)

	if block == nil {
		return nil, errors.New("Private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Create128bitKeyUsingMD5 function ...
func Create128bitKeyUsingMD5(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// ConvertUnsafeBytesToBase64String function ...
func ConvertUnsafeBytesToBase64String(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// ConvertBase64StringToUnsafeBytes function ...
func ConvertBase64StringToUnsafeBytes(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}

// EncryptOtp function ...
func EncryptOtp(otp, tenantHashID, mobileCountryCode, mobileNumber, salt string) (encryptedOtp string, cipherKey string, err error) {
	var cipherText []byte

	if cipherText, err = RSA_Encrypt([]byte(otp)); err != nil {
		return "", "", err
	}

	encryptedOtp = ConvertUnsafeBytesToBase64String(cipherText)
	return
}

// DecryptOtp function ...
func DecryptOtp(encryptedOtp, tenantHashID, mobileCountryCode, mobileNumber, salt string) (otp string, err error) {
	rawBytes, err := ConvertBase64StringToUnsafeBytes(encryptedOtp)

	if err != nil {
		return "", err
	}

	otpByte, err := RSA_Decrypt(rawBytes)
	otp = string(otpByte)
	return
}

// EncryptCompleteRSA function ...
func EncryptCompleteRSA(otp string) (encryptedOtp string, err error) {
	var cipherText []byte

	if cipherText, err = RSA_Encrypt([]byte(otp)); err != nil {
		return "", err
	}

	encryptedOtp = ConvertUnsafeBytesToBase64String(cipherText)
	return
}

// DecryptCompleteRSA function ...
func DecryptCompleteRSA(encryptedOtp string) (otp string, err error) {
	rawBytes, err := ConvertBase64StringToUnsafeBytes(encryptedOtp)

	if err != nil {
		return "", err
	}

	otpByte, err := RSA_Decrypt(rawBytes)
	otp = string(otpByte)
	return
}

// func EncryptPassword ...
func EncryptPassword(rawPassword string, email string, salt string) (encryptedPassword string, cipherKey string, err error) {

	var cipherText []byte
	cipherKey = Create128bitKeyUsingMD5(email + salt)

	if cipherText, err = AES_Encrypt(rawPassword, []byte(cipherKey)); err != nil {
		return "", "", err
	}

	encryptedPassword = ConvertUnsafeBytesToBase64String(cipherText)
	return
}

// func DecryptPassword ...
func DecryptPassword(encryptedPassword string, email string, salt string) (rawPassword string, err error) {

	var cipherText []byte

	if cipherText, err = ConvertBase64StringToUnsafeBytes(encryptedPassword); err != nil {
		return "", err
	}

	cipherKey := Create128bitKeyUsingMD5(email + salt)

	rawPassword, err = AES_Decrypt(cipherText, []byte(cipherKey))
	return
}
