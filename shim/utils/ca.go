package utils

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/tjfoc/gmsm/sm2"
	"strings"
)

func ParseSM2CertificateFromPemCert(pemBytes []byte) (*sm2.Certificate, error) {
	if len(pemBytes) == 0 {
		return nil, errors.New("invalid PEM. It must be different from nil")
	}
	block, pemBytes := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("pem certificate conversion error")
	}
	if block.Type != "CERTIFICATE" && block.Type != "" {
		return nil, errors.New("enrollment certificate should be a certificate")
	}
	return ParseSM2CertificateFromCert(block.Bytes)
}

func ParseSM2CertificateFromCert(bytes []byte) (*sm2.Certificate, error) {
	cert, err := sm2.ParseCertificate(bytes)
	if err != nil {
		return nil, err
	}
	key, _ := cert.PublicKey.(*ecdsa.PublicKey)
	switch key.Curve {
	case sm2.P256Sm2():
		return cert, nil
	default:
		return nil, errors.New("key's curve type error.")
	}
}

func ParseX509CertificateFromPemCert(pemBytes []byte) (*x509.Certificate, error) {
	if len(pemBytes) == 0 {
		return nil, errors.New("invalid PEM. It must be different from nil")
	}
	block, pemBytes := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("pem certificate conversion error")
	}
	if block.Type != "CERTIFICATE" && block.Type != "" {
		return nil, errors.New("enrollment certificate should be a certificate")
	}
	return ParseX509CertificateFromCert(block.Bytes)
}

func ParseX509CertificateFromCert(bytes []byte) (*x509.Certificate, error) {
	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, err
	}
	key, _ := cert.PublicKey.(*ecdsa.PublicKey)
	switch key.Curve {
	case elliptic.P256():
		return cert, nil
	default:
		return nil, errors.New("key's curve type error.")
	}
}

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
	if key, err := sm2.ParsePKCS8UnecryptedPrivateKey(der); err == nil {
		return key, nil
	}
	return nil, errors.New("tls: failed to parse private key")
}
