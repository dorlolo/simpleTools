package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func TestECB4(t *testing.T) {
	origtext := []byte("41e2188c94ee0:3205072020100901A01000:1643186261974")
	erytext := []byte("KaYhCwtWQ4KEifkL")
	//d := []byte("hello,ase")
	//key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(origtext))
	x1, err := encryptAES(origtext, erytext)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("加密后:", base64.StdEncoding.EncodeToString(x1))
	//x2, err := decryptAES(x1, erytext)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println("解密后:", string(x2))
}
