# 双向认证 ssl grpc


## ssl证书生成过程

### 1. 创建ca根证书

```bash
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 3650 -key ca.key -out ca.pem
```

### 2. 生成服务端证书

```bash
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
```

说明：需要依次输入国家，地区，城市，组织，组织单位，Common Name和Email。其中Common Name，可以写自己的名字或者域名，如果要支持https，Common Name应该与域名保持一致，否则会引起浏览器警告。
可以将证书发送给证书颁发机构（CA），CA验证过请求者的身份之后，会出具签名证书，需要花钱。另外，如果只是内部或者测试需求，也可以使用OpenSSL实现自签名。

### 3. 生成客户端证书

```bash
openssl ecparam -genkey -name secp384r1 -out client.key
openssl req -new -key client.key -out client.csr
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
```
