# ssl grpc通信

- tcp server
- http server


## ssl证书生成过程

### 1. 创建密钥

RSA私钥,需提供一个至少4位,最多1024位密码

```bash
openssl genrsa -des3 -out server.key 2048
```

### 2. 生成CSR (证书签名请求)

```bash
openssl req -new -key server.key -out server.csr
```

说明：需要依次输入国家，地区，城市，组织，组织单位，Common Name和Email。其中Common Name，可以写自己的名字或者域名，如果要支持https，Common Name应该与域名保持一致，否则会引起浏览器警告。
可以将证书发送给证书颁发机构（CA），CA验证过请求者的身份之后，会出具签名证书，需要花钱。另外，如果只是内部或者测试需求，也可以使用OpenSSL实现自签名。

### 3. 删除密钥中的密码

```bash
openssl rsa -in server.key -out no_password_server.key
```

### 4. 生成自签名证书

```bash
openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt
# 或无密码
openssl x509 -req -days 365 -in server.csr -signkey no_password_server.key -out no_password_server.crt
```

### 5. 生成pem格式的公钥

```bash
openssl x509 -in server.crt -out server.pem -outform PEM
# 或无密码
openssl x509 -in no_password_server.crt -out no_password_server.pem -outform PEM
```

