package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

var (
	err   error
	b     []byte
	block *pem.Block
	pub   ed25519.PublicKey
	priv  ed25519.PrivateKey
)

func main() {

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	fmt.Println("public = " + hex.EncodeToString(pub))
	fmt.Println("private = " + hex.EncodeToString(priv))

	b, err := x509.MarshalPKCS8PrivateKey(priv)

	if err != nil {
		fmt.Println(err)
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = os.WriteFile("private.pem", pem.EncodeToMemory(block), 0600)

	if err != nil {
		fmt.Println(err)
	}

	// public key
	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		fmt.Println(err)
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}

	err = os.WriteFile("public.pem", pem.EncodeToMemory(block), 0644)
}
