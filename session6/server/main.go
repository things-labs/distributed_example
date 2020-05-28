package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/thinkgos/grpcexample/session6/services"
)

func GotServerCrt() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../../cert/server.pem", "../../cert/server.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair失败 %v\n", err)
	}
	certPool := x509.NewCertPool()

	ca, _ := ioutil.ReadFile("../../cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequestClientCert,
		ClientCAs:    certPool,
	})
}

func gotClientCrt() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../../cert/client.pem", "../../cert/client.key")
	if err != nil {
		log.Fatalf("LoadX509KeyPair失败 %v\n", err)
	}
	certPool := x509.NewCertPool()

	ca, _ := ioutil.ReadFile("../../cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

}
func main() {
	rpcServer := grpc.NewServer(grpc.Creds(GotServerCrt()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	go rpcServer.Serve(listen)

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	rpcServer.ServeHTTP(w, r)
	//})
	//log.Fatal(http.ListenAndServeTLS(":8081", "../../cert/server.pem", "../../cert/server.key", nil))

	// http server grpc

	gmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(gotClientCrt())}
	err = services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", gmux)
}
