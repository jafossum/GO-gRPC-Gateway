package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/jafossum/go-grpc-gateway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	crt         = "../certificates/client.cer"
	key         = "../certificates/client.key"
	ca          = "../certificates/ca.cer"
)

func main() {

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	// Create the new TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ServerName:   "localhost", // NOTE: this is required!
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	// Connect to the gRPC Server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewYourServiceClient(conn)

	// Use default name or read form arguments
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// Contact the server and print out its response.
	ctx, channel := context.WithTimeout(context.Background(), time.Second)
	defer channel()
	r, err := c.Echo(ctx, &pb.StringMessage{Value: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Value)
}
