package gmt0018

import (
	"gmt0018/gmsm/sm2"
	"math/big"
)

type ECCrefPublicKey struct {
	Bits uint
	X    [64]byte
	Y    [64]byte
}
type ECCrefPrivateKey struct {
	Bits uint
	K    [64]byte
}

type ECCCipher struct {
	X [64]byte
	Y [64]byte
	M [32]byte
	L uint
	C [256]byte
}

type ECCSignature struct {
	R [64]byte
	S [64]byte
}

func (p *ECCrefPublicKey) ToPemPK() (pem string, err error) {
	sm2pk := sm2.PublicKey{}
	sm2pk.Curve = sm2.P256Sm2()
	sm2pk.X = new(big.Int).SetBytes(p.X[:])
	sm2pk.Y = new(big.Int).SetBytes(p.Y[:])

	der, err := sm2.WritePublicKeytoMem(&sm2pk, nil)
	pem = string(der)

	return
}

func (p *ECCrefPrivateKey) ToPemVK() (pem string, err error) {
	sm2vk := sm2.PrivateKey{}
	sm2vk.Curve = sm2.P256Sm2()
	sm2vk.D = new(big.Int).SetBytes(p.K[:])
	sm2vk.PublicKey = *CalculateSM2PK(p.K[:])
	der, err := sm2.WritePrivateKeytoMem(&sm2vk, nil)
	pem = string(der)
	return
}

func CalculateSM2PK(d []byte) *sm2.PublicKey {
	c := sm2.P256Sm2()
	pub := new(sm2.PublicKey)
	pub.Curve = c
	pub.X, pub.Y = c.ScalarBaseMult(d)
	return pub
}
