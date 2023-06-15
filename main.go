package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	randomString := string(b)

	encryptionKey := hex.EncodeToString([]byte(randomString))
	fmt.Println("Passphrase: " + randomString)
	fmt.Println("Encryption key: " + encryptionKey)
}
