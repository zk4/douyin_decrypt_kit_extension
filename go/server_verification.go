package main

/**
 * Copyright (C) 2019 ByteDance, Inc. All Rights Reserved.
 * Go 验证消息签名
 */

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func main() {
	// 替换成自己的 tpToken, 和收到的 timestamp, nonce, encrypt, msgSignature
	VerifyByteDanceServer("XXX",
		"XXX",
		"XXX",
		"XXX",
		"XXX")
}

func VerifyByteDanceServer(tpToken string, timestamp string, nonce string, encrypt string, msgSignature string) {
	values := []string{tpToken, timestamp, nonce, encrypt}
	sort.Strings(values)
	newMsgSignature := Sha1(strings.Join(values, ""))

	if newMsgSignature == msgSignature {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	encodeStr := fmt.Sprintf("%x", h.Sum(nil))
	return encodeStr
}
