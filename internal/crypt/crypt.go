package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	_rand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyHash(hash string, plainPwd []byte) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println("Incorrect password.")
		return false
	}

	return true
}

func GenerateKey(strength int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	length := strength

	var key strings.Builder

	for i := 0; i < length; i++ {
		key.WriteRune(chars[rand.Intn(len(chars))])
	}

	return key.String()
}

func EncryptKey(text, key string) string {
	toEncrypt := []byte(text)
	theKey := []byte(key)

	block, err := aes.NewCipher(theKey)
	if err != nil {
		log.Fatal(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(_rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	cipherText := aesGCM.Seal(nonce, nonce, toEncrypt, nil)
	return fmt.Sprintf("%x", cipherText)
}

func DecryptKey(encryptedText, key string) string {
	toDecrypt, _ := hex.DecodeString(encryptedText)
	theKey := []byte(key)

	block, err := aes.NewCipher(theKey)
	if err != nil {
		log.Fatal(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := toDecrypt[:nonceSize], toDecrypt[nonceSize:]

	decrypted, err := aesGCM.Open(nil, nonce, cipherText, nil)

	return fmt.Sprintf("%s", decrypted)
}
