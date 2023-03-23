package gmt0018

import "C"
import (
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

// DeviceHandle 设备句柄
type DeviceHandle struct{ P unsafe.Pointer }

// SessionHandle 会话句柄
type SessionHandle struct{ P unsafe.Pointer }

// KeyHandle 密钥句柄
type KeyHandle struct{ P unsafe.Pointer }

type DeviceInfo struct {
	// 设备生产厂商名称
	IssuerName string
	// 设备型号
	DeviceName string
	// 设备编号,包含:日期(8 学符)、批次号(3 字符)、流水号（5字符)，
	DeviceSerial string
	// 密码设备内部软件的版本号
	DeviceVersion string
	// 密码设备支持的接口规范版本号
	StandardVersion string
	// 前4字节表示支持的算法，表示方法为非对称算法标识按位或的结果;后4字节表示算法的最大模长，表示方法为支持的模长按位或的结果
	AsymAlgAbility string
	// 所有支持的对称算法,表示方法为对称算法标识按位或运算结果
	SymAlgAbility string
	// 所有支持的杂凑算法,表示方法为杂凑算法标识按位或运算结果
	HashAlgAbility string
	// 支持的最大文件存储空间(单位字节)
	BufferSize string
}

func (d *DeviceHandle) Get() string {
	return fmt.Sprintf("%p", d.P)
}
func (d *DeviceHandle) Set(str string) (err error) {
	if len(str) < 3 {
		err = errors.New("invalid ptr")
		return
	}
	ptr, err := strconv.ParseInt(str[2:], 16, 0)
	if err != nil {
		return err
	}
	d.P = (unsafe.Pointer(uintptr(ptr)))
	return
}

func (d *DeviceHandle) IsNil() bool {
	return d.P == nil
}

func (d *DeviceHandle) AssignNil() {
	d.P = nil
}

func (d *SessionHandle) Get() string {
	return fmt.Sprintf("%p", d.P)
}
func (d *SessionHandle) Set(str string) {
	ptr, _ := strconv.ParseInt(str[2:], 16, 0)
	d.P = (unsafe.Pointer(uintptr(ptr)))
}
func (d *SessionHandle) IsNil() bool {
	return d.P == nil
}

func (d *KeyHandle) Set(str string) (err error) {
	ptr, _ := strconv.ParseInt(str[2:], 16, 0)
	d.P = (unsafe.Pointer(uintptr(ptr)))
	return
}
func (d *KeyHandle) Get() string {
	return fmt.Sprintf("%p", d.P)
}

type SDF_HashInitArgs struct {
	Algid int
	PK    *ECCrefPublicKey
	UID   []byte
	//Ret   int
	Err error
}

type SDF_HashUpdateArgs struct {
	Data []byte
	//Ret  int
	Err error
}

type SDF_HashFinalArgs struct {
	Hash []byte
	//Ret  int
	Err error
}
