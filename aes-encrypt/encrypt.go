package main

import (
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
	"io"
	"io/ioutil"
)

func encrypt() {
	fmt.Println("Encryption program v0.0.1")

	text := []byte("My super secret")
	key := []byte("oneofstrongestpassphrase")

    // generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	// gcm or Galois/Counter Mode, is a mode of operation
    // for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	// creates a new byte array the size of the nonce
    // which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Non-encrypted text:", text)
	// here we encrypt our text using the Seal function
    // Seal encrypts and authenticates plaintext, authenticates the
    // additional data and appends the result to dst, returning the updated
    // slice. The nonce must be NonceSize() bytes long and unique for all
    // time, for a given key.
	fmt.Println(gcm.Seal(nonce, nonce, text, nil))

	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}
}