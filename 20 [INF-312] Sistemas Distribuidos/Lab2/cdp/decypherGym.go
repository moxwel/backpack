package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"l2/types"
)

func DesencriptarResultadoCombate(cifradoBase64 string, key_string string) (types.ResultadoCombate, error) {
	var res types.ResultadoCombate

	if len(key_string) != 32 {
		return res, fmt.Errorf("la clave debe tener 32 caracteres para AES256")
	}
	keyBytes := []byte(key_string)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return res, fmt.Errorf("error al crear el cifrador: %v", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cifradoBase64)
	if err != nil {
		return res, fmt.Errorf("error al decodificar base64: %v", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return res, fmt.Errorf("el texto cifrado es demasiado corto")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	err = json.Unmarshal(ciphertext, &res)
	if err != nil {
		return res, fmt.Errorf("error al deserializar el resultado de combate: %v", err)
	}

	return res, nil
}
