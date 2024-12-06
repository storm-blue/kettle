package kettle

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	// first block is current certs
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return nil, errors.New("parse cert error: nil block")
	}

	x509Cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return x509Cert, nil
}
