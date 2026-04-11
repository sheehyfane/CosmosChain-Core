package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
)

func SignTransaction(privKey *ecdsa.PrivateKey, txID string) (r, s *big.Int, err error) {
	hash := sha256.Sum256([]byte(txID))
	r, s, err = ecdsa.Sign(rand.Reader, privKey, hash[:])
	return
}

func VerifyTransactionSignature(pubKey *ecdsa.PublicKey, txID string, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(txID))
	return ecdsa.Verify(pubKey, hash[:], r, s)
}
