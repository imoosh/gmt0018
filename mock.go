package gmt0018

type mock struct{}

func (m mock) CreateFile(handle SessionHandle, fileName string, fileSize int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ReadFile(handle SessionHandle, fileName string) (content []byte, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) WriteFile(handle SessionHandle, fileName string, content []byte) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) DeleteFile(handle SessionHandle, fileName string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashInit(sessionHandle SessionHandle, algID int, pk *ECCrefPublicKey, id []byte) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashUpdate(sessionHandle SessionHandle, data []byte) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) HashFinal(sessionHandle SessionHandle) (hash []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) Encrypt(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (encdata []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) Decrypt(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (decdata []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CalculateMAC(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (mac []byte, outiv []byte, ret int, err error) {
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

func (m mock) ExternalVerify_ECC(sessionHandle SessionHandle, algID int, pk *ECCrefPublicKey, data []byte, sign *ECCSignature) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalSign_ECC(sessionHandle SessionHandle, iSKIndex int, data []byte) (sign *ECCSignature, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) InternalVerify_ECC(sessionHandle SessionHandle, iPKIndex int, data []byte, sign *ECCSignature) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExternalEncrypt_ECC(sessionHandle SessionHandle, pk *ECCrefPublicKey, data []byte) (eccCipher *ECCCipher, ret int, err error) {
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

func (m mock) ExportSignPublicKey_ECC(sessionHandle SessionHandle, index int) (pk *ECCrefPublicKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ExportEncPublicKey_ECC(sessionHandle SessionHandle, index int) (pk *ECCrefPublicKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyPair_ECC(sessionHandle SessionHandle, algID int, keybits int) (pk *ECCrefPublicKey, vk *ECCrefPrivateKey, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithIPK_ECC(sessionHandle SessionHandle, iPKIndex int, keybits int) (eccCipher *ECCCipher, keyhandle KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateKeyWithEPK_ECC(sessionHandle SessionHandle, algID int, keybits int, pk *ECCrefPublicKey) (eccCipher *ECCCipher, keyhandle KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ImportKeyWithISK_ECC(sessionHandle SessionHandle, iSKIndex int, eccCipher *ECCCipher) (keyhandle KeyHandle, ret int, err error) {
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

func (m mock) GenerateKeyWithKEK(sessionHandle SessionHandle, keybits int, aldID int, kekIndex int) (key []byte, keyhandle KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ImportKeyWithKEK(sessionHandle SessionHandle, aldID int, kekIndex int, key []byte) (keyhandle KeyHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) DestroyKey(sessionHandle SessionHandle, keyhandle KeyHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) OpenDevice() (handle DeviceHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CloseDevice(handle DeviceHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) OpenSession(devHandle DeviceHandle) (sesHandle SessionHandle, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) CloseSession(handle SessionHandle) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GetDeviceInfo(handle SessionHandle) (info *DeviceInfo, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GenerateRandom(handle SessionHandle, len int) (rand []byte, ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) GetPrivateKeyAccessRight(handle SessionHandle, index int, pwd string) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}

func (m mock) ReleasePrivateKeyAccessRight(handle SessionHandle, index int) (ret int, err error) {
	//TODO implement me
	panic("implement me")
}
