package main

import (
	"crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
)

func main() {
	fmt.Println("Decryption program v0.0.1")

	key := []byte("oneofstrongestpassphrase")
	ciphertext, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext :=ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}