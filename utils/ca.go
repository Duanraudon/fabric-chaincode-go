package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	x509sm2 "github.com/tjfoc/gmsm/x509"
	"strings"
)

func ParseSM2PrivateKey(pemBytes []byte) (*sm2.PrivateKey, bool) {
	if len(pemBytes) == 0 {
		return nil, false
	}
	block, pemBytes := pem.Decode(pemBytes)
	if block == nil {
		return nil, false
	}
	if block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, " PRIVATE KEY") {
		if priKey, err := parsePrivateKey(block.Bytes); err == nil {
			if priKey, ok := priKey.(*sm2.PrivateKey); ok {
				return priKey, true
			}
		}
	}
	return nil, false
}

func ParseECDSAPrivateKey(pemBytes []byte) (*ecdsa.PrivateKey, bool) {
	if len(pemBytes) == 0 {
		return nil, false
	}
	block, pemBytes := pem.Decode(pemBytes)
	if block == nil {
		return nil, false
	}
	if block.Type == "PRIVATE KEY" || strings.HasSuffix(block.Type, " PRIVATE KEY") {
		if priKey, err := parsePrivateKey(block.Bytes); err == nil {
			if priKey, ok := priKey.(*ecdsa.PrivateKey); ok {
				return priKey, true
			}
		}
	}
	return nil, false
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey:
			return key, nil
		default:
			return nil, errors.New("tls: found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err := x509sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, errors.New("tls: failed to parse private key")
}
