package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

func getCipher(key string) cipher.Block {
	encryptKeyNotEmpty(key)
	cipher, _ := aes.NewCipher(generateKey([]byte(key)))
	return cipher
}

func encryptKeyNotEmpty(key string) {
	if key == "" {
		panic("aes key is empty")
	}
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
