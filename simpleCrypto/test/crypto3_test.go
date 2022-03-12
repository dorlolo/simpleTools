/*
 * @FileName:   crypto_test3.go
 * @Author:		JunjieXu
 * @CreateTime:	2022/1/26 下午5:09
 * @Description:
 */
package test

import (
	"bytes"
	"crypto/aes"
	"crypto/des"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"testing"
)

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func getSha1PrngKey(keyword string) (key []byte) {
	data := []byte(keyword)
	hashs := SHA1(SHA1(data))
	key = hashs[0:8]
	return key
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	fmt.Println(padding)
	fmt.Println(len(ciphertext))
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()

	//src = ZeroPadding(src, bs)
	src = PKCS5Padding(src, bs)
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func TestECB2(t *testing.T) {
	origtext := []byte("41e2188c94ee0:3205072020100901A01000:1643186261974")
	erytext, err := DesEncrypt(origtext, []byte("KaYhCwtWQ4KEifkL"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(erytext))
	//
	//s, _ := DesDecrypt(erytext, getSha1PrngKey("123456"))
	//fmt.Println(string(s))
}
