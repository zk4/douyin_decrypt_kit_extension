<?php
/**
 * Copyright (C) 2019 ByteDance, Inc. All Rights Reserved.
 * PHP AES 消息加解密
 */
class AesEncryptUtil
{
    public static $blockSize = 32;
    //your aesKey
    private $aesKey;
    /**
     * AesEncryptNewUtilUtil constructor.
     * @param $encodingAesKey
     */
    public function __construct($encodingAesKey)
    {
        $this->aesKey = base64_decode($encodingAesKey . "=");
    }

    /**
     * 对密文进行解密
     * @param $encrypted
     * @return bool|string
     * @throws Exception
     */
    public function decrypt($encrypted)
    {
        try {
            // 使用BASE64对需要解密的字符串进行解码
            $ciphertextDec = base64_decode($encrypted);
            $iv = substr($this->aesKey, 0, 16);
            // 解密
            $decrypted = openssl_decrypt($ciphertextDec, 'aes-256-cbc', $this->aesKey, OPENSSL_RAW_DATA | OPENSSL_NO_PADDING, $iv);
        } catch (Exception $e) {
            throw new Exception('AesEncryptUtil AES解密串非法，小于16位;');
        }
        try {
            // 去除补位字符
            $result = $this->decode($decrypted);
            // 去除16位随机字符串,网络字节序和 tp appid
            if (strlen($result) < 16) {
                throw new Exception('AesEncryptUtil AES解密串非法，小于16位;');
            }
            // 去除16位随机字符串
            $content = substr($result, 16, strlen($result));

            // 获取消息体长度
            $lenList = unpack("N", substr($content, 16, 4));
            $postBodyLen = $lenList[1];

            // 获取消息体
            $postBodyMsg = substr($content, 20, $postBodyLen);

            // 获取消息体的第三方平台 appid
            $fromTpAppId = substr($content, 20 + $postBodyLen);
        } catch (Exception $e) {
            throw new Exception('AesEncryptUtil AES解密串非法，小于16位;');
        }
        return $postBodyMsg;
    }


    /**
     * 对解密后的明文进行补位删除
     * @param $text
     * @return bool|string
     */
    private function decode($text)
    {
        $pad = ord(substr($text, -1));
        if ($pad < 1 || $pad > self::$blockSize) {
            $pad = 0;
        }
        return substr($text, 0, (strlen($text) - $pad));
    }

}

// 替换成自己的 aesKey
$test = new AesEncryptUtil("XXX");
// 替换成收到的 encrypt
echo $test -> decrypt("XXX");