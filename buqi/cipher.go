package buqi

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
)

// Password 密码类型 
type Password [256]byte

// Cipher 密码编码解码
type Cipher struct {
	// 编码
	encode *Password
	// 解码
	decode *Password
}

// String 密码转字符串
func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

// ParsePassword 解析密码
func ParsePassword(str string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(str))
	if err != nil || len(bs) != 256 {
		return nil, errors.New("密码格式有误")
	}
	password := Password{}
	copy(password[:], bs)
	bs = nil
	return &password, nil
}

// RandPassword 随机密码
func RandPassword() *Password {
	// 生成随机数组
	arr := rand.Perm(256)
	pwd := &Password{}
	for k, v := range arr {
		pwd[k] = byte(v)
		// 若键值相同则进行递归
		if k == v {
			return RandPassword()
		}
	}
	return pwd
}

// encrypt 加密数据
func (c *Cipher) encrypt(b []byte) {
	for i, v := range b {
		b[i] = c.encode[v]
	}
}

// decrypt 解密数据
func (c *Cipher) decrypt(b []byte) {
	for i, v := range b {
		b[i] = c.decode[v]
	}
}

// NewCipher 新建一个编码解码器
func NewCipher(encodePassword *Password) *Cipher {
	decodePassword := &Password{}
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}
	return &Cipher{
		encode: encodePassword,
		decode: decodePassword,
	}
}
