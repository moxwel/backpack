package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"l2/types"
)

func EncriptarResultadoCombate(resultado types.ResultadoCombate, key_string string) (string, error) {
	if len(key_string) != 32 {
		return "", fmt.Errorf("la clave debe tener 32 caracteres para AES256")
	}
	keyBytes := []byte(key_string)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("error al crear el cifrador: %v", err)
	}

	jsonData, err := json.Marshal(resultado)
	if err != nil {
		return "", fmt.Errorf("error al serializar el resultado de combate: %v", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(jsonData))
	iv := ciphertext[:aes.BlockSize]
	reader := bufio.NewReader(rand.Reader)
	if _, err := reader.Read(iv); err != nil {
		return "", fmt.Errorf("error al generar el IV: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], jsonData)

	base64Ciphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return base64Ciphertext, nil
}
