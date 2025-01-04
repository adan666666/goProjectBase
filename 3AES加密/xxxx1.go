package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var KEY = []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab") //32
var IV = []byte("aaaaaaaaaaaaaaaa")

func main() {
	plaintext := []byte("123456")
	// 加密
	encrypted, err := AESEncrypt1(plaintext, KEY, IV)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Printf("Encrypted: %x\n", encrypted)

}

func PKCS7Padding1(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// AESEncrypt 使用AES进行加密 PKCS7
func AESEncrypt1(plaintext, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding1(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}
