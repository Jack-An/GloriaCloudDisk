package common

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

type (
	Crypt interface {
		Encrypt(str string) string
		Verify(secret string, hash string) bool
	}

	CryptContext struct {
		Schema string
	}
)

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func (c *CryptContext) Encrypt(secret string) (string, error) {
	src := []byte(secret)
	block, err := des.NewCipher([]byte(c.Schema))
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func (c *CryptContext) Verify(secret string, hash string) (bool, error) {
	src, err := hex.DecodeString(hash)
	if err != nil {
		return false, err
	}
	block, err := des.NewCipher([]byte(c.Schema))
	if err != nil {
		return false, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return false, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out) == secret, nil
}
