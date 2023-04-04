提供验证消息签名以及解密消息的模版，有 PHP、JAVA、GO 三种语言。
验证消息签名：
    java 文件夹下的 ServerVerification 文件中的 verify 方法。
    php 文件夹下的 serverVerification 文件中的 verify 方法。
    go 文件夹下的 server_verification 文件中的 VerifyByteDanceServer 方法。
解密消息：
    java 文件夹下的 MsgDecrypt 文件中的 decrypt 方法，需要配置 pom.xml 文件。
    php 文件夹下的 msgDecrypt 文件中的 decrypt 方法。
    go 文件夹下的 msg_decrypt 文件中的 DecryptMsg 方法。


---
扩展了两个语言的版本，python 与 node， node 可以方便的集成到 postman 里调试。 
