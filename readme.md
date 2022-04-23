# EnDecrypt使用教程

## 1.安装

```shell
#1.解压文件夹
$ cd endecrypt
#2.查看文件
$ ls
cmd  endecrypt.go  go.mod  go.sum  imp readme.md
#3.调整
$ go mod tidy
#4.编译
$ go build -o endecrypt endecrypt.go
```

## 2.测试

### 2.1帮助命令

```shell
#1.查看帮助
$ ./endecrypt --help
The tool for encryption & decryption with common algorithm.

Usage:
  endecrypt [flags]
  endecrypt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     The command of decryption. #用于解密
  encrypt     The command of encryption. #用于加密
  help        Help about any command     #查看帮助信息
  version     Print endecrypt version info. #打印软件版本

Flags:
  -h, --help             help for endecrypt
  -V, --version string   The version of endecrypt. (default "1.0")

Use "endecrypt [command] --help" for more information about a command.
#2.查看加密帮助
$ ./endecrypt encrypt --help
The command of encryption.

Usage:
  endecrypt encrypt [flags]

Flags:
  -c, --eccpublic string   The eccpublic key file to be encryted. (default "eccpublic.pem")#指定ecc公钥文件
  -h, --help               help for encrypt                  
  -k, --key string         The key to be encryted. (default "1234567887654321") #密钥
  -p, --public string      The public key file to be encryted. (default "public.pem") #指定rsa算法公钥文件
  -e, --text string        The text to be encryted.          #明文
  -t, --type string        The type of encryption[aescbc/aesctr/des/3des/sha256hex/sha512hex/ecc/rsa]  (default "aescbc") #加密选项
#3.查看解密帮助
$ ./endecrypt decrypt --help
The command of decryption.

Usage:
  endecrypt decrypt [flags]

Flags:
  -C, --eccprivate string   The eccprivate key file to be decryted. (default "eccprivate.pem")#ecc解密私钥文件
  -h, --help                help for decrypt
  -k, --key string          The key to be decryted. (default "1234567887654321") #密钥
  -d, --password string     The text to be decryted.  #解密的密文
  -p, --private string      The private key file to be decryted. (default "private.pem")#rsa解密私钥
  -t, --type string         The type of decryption[aescbc/aesctr/des/3des/ecc/rsa] (default "aescbc")#解密选项
```

### 2.2 加密实践

```shell
#aes算法 cbc模式
$./endecrypt encrypt -t aescbc -e hello
PRvjGawvUcZ5pvbYVNE4Sw==
#aes算法 ctr模式
$ ./endecrypt encrypt -t aesctr -e hello
/cscDZE=
# des算法
$ ./endecrypt encrypt -t des -e hello
IXQccBtyDRk=
# 3des算法
$ ./endecrypt encrypt -t 3des -e hello
UHdm1ehDF4k=
# sha256hex算法sha512hex算法
$ ./endecrypt encrypt -t sha256hex -e hello
2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
$ ./endecrypt encrypt -t sha512hex -e hello
9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043

#生成ecc公钥和私钥文件
$ ./endecrypt encrypt -t ecc

$ ls
cmd  eccprivate.pem  eccpublic.pem  endecrypt  endecrypt.go  go.mod  go.sum  imp

$ ./endecrypt encrypt -t ecc -e hello
04c77ff64d5fb70a2769beeab1d8bbafb5197ffc6fb0d5d65d16054612476d152c582f3d2c4acaa686b848a5368a734638c66ed440f449d6d9609c59f1121ddd8faaa75f36d021473605979d10e710951a59f918423f4ed30d16cf68202db7f7ba1b206109a64487b663f30b548997425bbeb065747a
#生成rsa公钥和私钥文件
$ ./endecrypt encrypt -t rsa

$ ls
cmd  eccprivate.pem  eccpublic.pem  endecrypt  endecrypt.go  go.mod  go.sum  imp  private.pem  public.pem
$ ./endecrypt encrypt -t rsa -e hello
2ffd58249b37f3d063ac418c90163ba45e777bf226dbea049689f3eaec037b4df1d152de288513a1a57ad99bf435770a330e5efa98a840f6d1420d79c34035b54b2e3cb78dca29622e2c23735512a14bf71ba612983741e0988e0dfb996e9221513b46fdc9b28f8f08300c7d29a71f570a46a71c470e94b67d7a9b9365ae12549fe4cb1450ef6d084607dd8e6ddaf1327dccdd2f752a0c696547cc90cfc922c33bf1dfb2f5eb825eb01c50b503da98c794e8dd33da59905fc5bf63ce9f00d4b0f294cfe1e470cb3b81e8dbf294fa5aed8d88864e0f55bde6539e22bf6c88c8fdc1a7533cf157034e73f11491343fb20c623e0a32653d7823b482b18277bf8161
```

### 2.3 解密实践

```shell
#aes算法 cbc模式
$ ./endecrypt decrypt -t aescbc -d PRvjGawvUcZ5pvbYVNE4Sw==
hello
#aes算法 ctr模式
$ ./endecrypt decrypt -t aesctr -d /cscDZE=
hello
#3des算法
$ ./endecrypt decrypt -t 3des -d UHdm1ehDF4k=
hello
#des算法
$ ./endecrypt decrypt -t des -d  IXQccBtyDRk=
hello
#rsa算法
$./endecrypt decrypt -t rsa -d 2ffd58249b37f3d063ac418c90163ba45e777bf226dbea049689f3eaec037b4df1d152de288513a1a57ad99bf435770a330e5efa98a840f6d1420d79c34035b54b2e3cb78dca29622e2c23735512a14bf71ba612983741e0988e0dfb996e9221513b46fdc9b28f8f08300c7d29a71f570a46a71c470e94b67d7a9b9365ae12549fe4cb1450ef6d084607dd8e6ddaf1327dccdd2f752a0c696547cc90cfc922c33bf1dfb2f5eb825eb01c50b503da98c794e8dd33da59905fc5bf63ce9f00d4b0f294cfe1e470cb3b81e8dbf294fa5aed8d88864e0f55bde6539e22bf6c88c8fdc1a7533cf157034e73f11491343fb20c623e0a32653d7823b482b18277bf8161
hello
#ecc算法
$ ./endecrypt decrypt -t ecc -d 04c77ff64d5fb70a2769beeab1d8bbafb5197ffc6fb0d5d65d16054612476d152c582f3d2c4acaa686b848a5368a734638c66ed440f449d6d9609c59f1121ddd8faaa75f36d021473605979d10e710951a59f918423f4ed30d16cf68202db7f7ba1b206109a64487b663f30b548997425bbeb065747a
hello
```