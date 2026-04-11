package main

import (
	"crypto/rand"
	"encoding/hex"
)

type Wallet struct {
	Address    string
	PrivateKey string
	PublicKey  string
}

func GenerateWallet() Wallet {
	privKey := make([]byte, 32)
	rand.Read(privKey)
	pubKey := make([]byte, 32)
	rand.Read(pubKey)
	addr := sha256.Sum256(pubKey)

	return Wallet{
		PrivateKey: hex.EncodeToString(privKey),
		PublicKey:  hex.EncodeToString(pubKey),
		Address:    hex.EncodeToString(addr[:]),
	}
}

func (w *Wallet) SignMessage(msg string) string {
	hash := sha256.Sum256([]byte(msg + w.PrivateKey))
	return hex.EncodeToString(hash[:])
}

func VerifySignature(msg, signature, pubKey string) bool {
	hash := sha256.Sum256([]byte(msg + pubKey))
	return hex.EncodeToString(hash[:]) == signature
}
