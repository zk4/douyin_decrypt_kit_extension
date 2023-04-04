package main

/**
 * Copyright (C) 2019 ByteDance, Inc. All Rights Reserved.
 * Go AES 消息加解密
 */

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	// 替换成自己的 encodeAesKey, 和收到的消息体
	DecryptMsg("XXX",
		"XXX")
}

func DecryptMsg(encodeAesKey string, encryptMsg string) {
	// get aes key
	AESKey, _ := base64.StdEncoding.DecodeString(encodeAesKey + "=")

	// decrypt msg
	decryptMsg, _ := Decrypt(encryptMsg, string(AESKey))

	// plain text
	plainText := []byte(decryptMsg)
	buf := bytes.NewBuffer(plainText[16:20])
	var length int32
	_ = binary.Read(buf, binary.BigEndian, &length)

	// 推送的第三方 AppID
	appIDStart := 20 + length
	tpAppId := string(plainText[appIDStart:])
	fmt.Printf("thirdparty appid: %s\n", tpAppId)

	// 获取正常的消息体
	msgBody := string(plainText[20 : 20+length])
	fmt.Printf("decode msg body: %s\n", msgBody)

	// 返回解析的消息 json 串
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(msgBody), &result)
	fmt.Printf("msg %+v", result)
}

func Decrypt(rawData, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	if err != nil {
		return "", err
	}
	dnData, err := AESCBCDecrypt(data, []byte(key))
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

func AESCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		return []byte{}, errors.New("cipherText too short")
	}

	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]
	if len(encryptData)%blockSize != 0 {
		return []byte{}, errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	encryptData = PKCS7UnPadding(encryptData)

	return encryptData, nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}