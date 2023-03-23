package sm2

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/gob"
)

func init() {
	gob.Register(sm2P256Curve{})
	gob.Register(x509.CertificateRequest{})
	gob.Register(CertificateRequest{})
	gob.Register(ecdsa.PublicKey{})
}
