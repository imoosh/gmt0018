package test

import (
	"gitea.com/amchd/SSN_gmt0018"
)

type mock struct{}

func (m mock) CreateFile(handle gmt0018.SessionHandle, fileName string, fileSize int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ReadFile(handle gmt0018.SessionHandle, fileName string) (content []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) WriteFile(handle gmt0018.SessionHandle, fileName string, content []byte) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) DeleteFile(handle gmt0018.SessionHandle, fileName string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashInit(sessionHandle gmt0018.SessionHandle, algID int, pk *gmt0018.ECCrefPublicKey, id []byte) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashUpdate(sessionHandle gmt0018.SessionHandle, data []byte) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashFinal(sessionHandle gmt0018.SessionHandle) (hash []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) Encrypt(handle gmt0018.SessionHandle, keyhandle gmt0018.KeyHandle, algID int, iv []byte, data []byte) (encdata []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) Decrypt(handle gmt0018.SessionHandle, keyhandle gmt0018.KeyHandle, algID int, iv []byte, data []byte) (decdata []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CalculateMAC(handle gmt0018.SessionHandle, keyhandle gmt0018.KeyHandle, algID int, iv []byte, data []byte) (mac []byte, outiv []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExternalPublicKeyOperation_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalPublicKeyOperation_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalPrivateKeyOperation_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExternalVerify_ECC(sessionHandle gmt0018.SessionHandle, algID int, pk *gmt0018.ECCrefPublicKey, data []byte, sign *gmt0018.ECCSignature) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalSign_ECC(sessionHandle gmt0018.SessionHandle, iSKIndex int, data []byte) (sign *gmt0018.ECCSignature, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalVerify_ECC(sessionHandle gmt0018.SessionHandle, iPKIndex int, data []byte, sign *gmt0018.ECCSignature) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExternalEncrypt_ECC(sessionHandle gmt0018.SessionHandle, pk *gmt0018.ECCrefPublicKey, data []byte) (eccCipher *gmt0018.ECCCipher, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExportSignPublicKey_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExportEncPublicKey_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyPair_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithIPK_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithEPK_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ImportKeyWithISK_RSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExchangeDigitEnvelopeBaseOnRSA() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExportSignPublicKey_ECC(sessionHandle gmt0018.SessionHandle, index int) (pk *gmt0018.ECCrefPublicKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExportEncPublicKey_ECC(sessionHandle gmt0018.SessionHandle, index int) (pk *gmt0018.ECCrefPublicKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyPair_ECC(sessionHandle gmt0018.SessionHandle, algID int, keybits int) (pk *gmt0018.ECCrefPublicKey, vk *gmt0018.ECCrefPrivateKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithIPK_ECC(sessionHandle gmt0018.SessionHandle, iPKIndex int, keybits int) (eccCipher *gmt0018.ECCCipher, keyhandle gmt0018.KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithEPK_ECC(sessionHandle gmt0018.SessionHandle, algID int, keybits int, pk *gmt0018.ECCrefPublicKey) (eccCipher *gmt0018.ECCCipher, keyhandle gmt0018.KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ImportKeyWithISK_ECC(sessionHandle gmt0018.SessionHandle, iSKIndex int, eccCipher *gmt0018.ECCCipher) (keyhandle gmt0018.KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateAgreementDataWithECC() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithECC() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateAgreementDataAndKeyWithECC() {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExchangeDigitEnvelopeBaseOnECC() {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithKEK(sessionHandle gmt0018.SessionHandle, keybits int, aldID int, kekIndex int) (key []byte, keyhandle gmt0018.KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ImportKeyWithKEK(sessionHandle gmt0018.SessionHandle, aldID int, kekIndex int, key []byte) (keyhandle gmt0018.KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) DestroyKey(sessionHandle gmt0018.SessionHandle, keyhandle gmt0018.KeyHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) OpenDevice() (handle gmt0018.DeviceHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CloseDevice(handle gmt0018.DeviceHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) OpenSession(devHandle gmt0018.DeviceHandle) (sesHandle gmt0018.SessionHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CloseSession(handle gmt0018.SessionHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GetDeviceInfo(handle gmt0018.SessionHandle) (info *gmt0018.DeviceInfo, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateRandom(handle gmt0018.SessionHandle, len int) (rand []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GetPrivateKeyAccessRight(handle gmt0018.SessionHandle, index int, pwd string) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ReleasePrivateKeyAccessRight(handle gmt0018.SessionHandle, index int) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}
