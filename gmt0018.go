package SSN_gmt0018

import "C"

// GMT 0018-2012 密码设备应用接口规范

type GMT0018Handler interface {
	// DeviceManager 设备管理接口
	DeviceManager

	// KeyManager 密钥管理接口
	KeyManager

	// AsymmetricAlgorithmsComputer 非对称算法运算类函数
	AsymmetricAlgorithmsComputer

	// SymmetricAlgorithmsComputer 对称算法运算类函数
	SymmetricAlgorithmsComputer

	// HashAlgorithmsComputer 杂凑运算类函数
	HashAlgorithmsComputer

	// UserFileOperator 用户文件操作类函数
	UserFileOperator
}

// DeviceManager 设备管理接口
type DeviceManager interface {
	// OpenDevice 打开设备
	OpenDevice() (handle DeviceHandle, ret int, err error)

	// CloseDevice 关闭设备
	CloseDevice(handle DeviceHandle) (ret int, err error)

	// OpenSession 创建会话
	OpenSession(devHandle DeviceHandle) (sesHandle SessionHandle, ret int, err error)

	// CloseSession 关闭会话
	CloseSession(handle SessionHandle) (ret int, err error)

	// GetDeviceInfo 获取设备信息
	GetDeviceInfo(handle SessionHandle) (info *DeviceInfo, ret int, err error)

	// GenerateRandom 产生随机数
	GenerateRandom(handle SessionHandle, len int) (rand []byte, ret int, err error)

	// GetPrivateKeyAccessRight 获取私钥使用权限
	GetPrivateKeyAccessRight(handle SessionHandle, index int, pwd string) (ret int, err error)

	// ReleasePrivateKeyAccessRight 释放私钥使用权限
	ReleasePrivateKeyAccessRight(handle SessionHandle, index int) (ret int, err error)
}

// KeyManager 密钥管理接口
type KeyManager interface {
	// ExportSignPublicKey_RSA 导出 RSA 签名公钥
	ExportSignPublicKey_RSA()

	// ExportEncPublicKey_RSA 导出 RSA 加密公钥
	ExportEncPublicKey_RSA()

	// GenerateKeyPair_RSA 产生 RSA 非对称密钥对并输出
	GenerateKeyPair_RSA()

	// GenerateKeyWithIPK_RSA 生成会话密钥并用内部 RSA 公钥加密输出
	GenerateKeyWithIPK_RSA()

	// GenerateKeyWithEPK_RSA 生成会话密钥并用外部 RSA 公钥加密输出
	GenerateKeyWithEPK_RSA()

	// ImportKeyWithISK_RSA 导人会话密钥并用内部 RSA 私钥解密
	ImportKeyWithISK_RSA()

	// ExchangeDigitEnvelopeBaseOnRSA 基于 RSA 算法的数字信封转换
	ExchangeDigitEnvelopeBaseOnRSA()

	// ExportSignPublicKey_ECC 导出 ECC 签名公钥
	ExportSignPublicKey_ECC(sessionHandle SessionHandle, index int) (pk *ECCrefPublicKey, ret int, err error)

	// ExportEncPublicKey_ECC 导出 ECC 加密公钥
	ExportEncPublicKey_ECC(sessionHandle SessionHandle, index int) (pk *ECCrefPublicKey, ret int, err error)

	// GenerateKeyPair_ECC 产生 ECC 非对称密钥对并输出
	GenerateKeyPair_ECC(sessionHandle SessionHandle, algID int, keybits int) (pk *ECCrefPublicKey, vk *ECCrefPrivateKey, ret int, err error)

	// GenerateKeyWithIPK_ECC 生成会话密钥并用内部 ECC 公钥加密输出
	GenerateKeyWithIPK_ECC(sessionHandle SessionHandle, iPKIndex int, keybits int) (eccCipher *ECCCipher, keyhandle KeyHandle, ret int, err error)

	// GenerateKeyWithEPK_ECC 生成会话密钥并用外部 ECC 公钥加密输出
	GenerateKeyWithEPK_ECC(sessionHandle SessionHandle, algID int, keybits int, pk *ECCrefPublicKey) (eccCipher *ECCCipher, keyhandle KeyHandle, ret int, err error)

	// ImportKeyWithISK_ECC 导入会话密钥并用内部 ECC 私钥解密
	ImportKeyWithISK_ECC(sessionHandle SessionHandle, iSKIndex int, eccCipher *ECCCipher) (keyhandle KeyHandle, ret int, err error)

	// GenerateAgreementDataWithECC 生成密钥协商参数并输出
	GenerateAgreementDataWithECC()

	// GenerateKeyWithECC 计算会话密钥
	GenerateKeyWithECC()

	// GenerateAgreementDataAndKeyWithECC 产生协商数据并计算会话密钥
	GenerateAgreementDataAndKeyWithECC()

	// ExchangeDigitEnvelopeBaseOnECC 基于 ECC 算法的数字信封转换
	ExchangeDigitEnvelopeBaseOnECC()

	// GenerateKeyWithKEK 生成会话密钥并用密钥加密密钥加密输出
	GenerateKeyWithKEK(sessionHandle SessionHandle, keybits int, aldID int, kekIndex int) (key []byte, keyhandle KeyHandle, ret int, err error)

	// ImportKeyWithKEK 导入会话密钥并用密钥加密密钥解密
	ImportKeyWithKEK(sessionHandle SessionHandle, aldID int, kekIndex int, key []byte) (keyhandle KeyHandle, ret int, err error)

	// DestroyKey 销毁会话密钥
	DestroyKey(sessionHandle SessionHandle, keyhandle KeyHandle) (ret int, err error)
}

// AsymmetricAlgorithmsComputer 非对称算法运算类函数
type AsymmetricAlgorithmsComputer interface {
	// ExternalPublicKeyOperation_RSA 外部公钥 RSA 运算
	ExternalPublicKeyOperation_RSA()

	// InternalPublicKeyOperation_RSA 内部公钥 RSA 运算
	InternalPublicKeyOperation_RSA()

	// InternalPrivateKeyOperation_RSA 内部私钥 RSA 运算
	InternalPrivateKeyOperation_RSA()

	// ExternalVerify_ECC 外部密钥 ECC 验证
	ExternalVerify_ECC(sessionHandle SessionHandle, algID int, pk *ECCrefPublicKey, data []byte, sign *ECCSignature) (ret int, err error)

	// InternalSign_ECC 内部密钥 ECC 签名
	InternalSign_ECC(sessionHandle SessionHandle, iSKIndex int, data []byte) (sign *ECCSignature, ret int, err error)

	// InternalVerify_ECC 内部密钥 ECC 验证
	InternalVerify_ECC(sessionHandle SessionHandle, iPKIndex int, data []byte, sign *ECCSignature) (ret int, err error)

	// ExternalEncrypt_ECC 外部密钥 ECC 加密
	ExternalEncrypt_ECC(sessionHandle SessionHandle, pk *ECCrefPublicKey, data []byte) (eccCipher *ECCCipher, ret int, err error)
}

// SymmetricAlgorithmsComputer 对称算法运算类函数
type SymmetricAlgorithmsComputer interface {
	// Encrypt 对称加密
	Encrypt(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (encdata []byte, ret int, err error)

	// Decrypt 对称解密
	Decrypt(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (decdata []byte, ret int, err error)

	// CalculateMAC 计算MAC
	CalculateMAC(handle SessionHandle, keyhandle KeyHandle, algID int, iv []byte, data []byte) (mac []byte, outiv []byte, ret int, err error)
}

// HashAlgorithmsComputer 杂凑运算类函数
type HashAlgorithmsComputer interface {
	// HashInit 杂凑运算初始化
	// @handle:     会话句柄
	// @algID:      签名者公钥
	// @publicKey:  签名者公钥，当algID为SGD_SM3时有效
	// @id:         签名者id，当algID为SGD_SM3时有效
	HashInit(sessionHandle SessionHandle, algID int, pk *ECCrefPublicKey, id []byte) (ret int, err error)

	// HashUpdate 多包杂凑运算
	// @handle:     会话句柄
	// @data:       数据明文
	HashUpdate(sessionHandle SessionHandle, data []byte) (ret int, err error)

	// HashFinal 杂凑运算结束
	// @handle:     会话句柄
	HashFinal(sessionHandle SessionHandle) (hash []byte, ret int, err error)
}

// UserFileOperator 用户文件操作类函数
type UserFileOperator interface {
	// CreateFile 创建文件
	CreateFile(handle SessionHandle, fileName string, fileSize int) (err error)

	// ReadFile 读文件
	ReadFile(handle SessionHandle, fileName string) (content []byte, err error)

	// WriteFile 写文件
	WriteFile(handle SessionHandle, fileName string, content []byte) (err error)

	// DeleteFile 删除文件
	DeleteFile(handle SessionHandle, fileName string) (err error)
}
