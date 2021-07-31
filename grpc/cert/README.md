# ssl证书生成过程

## CA证书

### 1. 生成CA密钥
```bash
openssl genrsa -out ca.key 2048
```

### 2. 生成CA根证书

```bash
openssl req -sha256 -new -x509 \
      -days 3650 \
      -key ca.key \
      -out ca.crt \
      -subj "/C=CN/ST=FJ/L=FZ/O=didong/OU=thinkgo/CN=thinkgos.cn"
```

## 服务端证书

### 1. 生成服务器密钥

```bash
openssl genrsa -out server.key 2048
```

### 2. 生成服务器证书请求文件

```bash
openssl req -new -sha256 \
    -key server.key \
    -subj "/C=CN/ST=FJ/L=FZ/O=didong/OU=thinkgo/CN=thinkgos.cn" \
    -reqexts SAN \
    -config <(cat /usr/lib/ssl/openssl.cnf <(printf "[SAN]\nsubjectAltName=DNS:*.thinkgos.cn")) \
    -out server.csr
```

### 3. CA签发服务器证书
```bash
openssl x509 -req \
    -days 3650 \
   -in server.csr \
   -CA ca.crt \
   -CAkey ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /usr/lib/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.thinkgos.cn")) \
   -out server.crt
```

## 客户端证书

### 1. 生成客户端密钥

```bash
openssl genrsa -out client.key 2048
```

### 2. 生成客户端证书请求文件

```bash
openssl req -new \
    -key client.key \
	-subj "/C=CN/ST=FJ/L=FZ/O=didong/OU=thinkgo/CN=thinkgos.cn" \
	-reqexts SAN \
	-config <(cat /usr/lib/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.thinkgos.cn")) \
    -out client.csr
```

### 3. CA签发客户端证书
```bash
openssl x509 -req \
    -days 3650 \
   -in client.csr \
   -CA ca.crt \
   -CAkey ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat /usr/lib/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.thinkgos.cn")) \
  -out client.crt
```