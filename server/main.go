package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	pb "github.com/jafossum/go-grpc-gateway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
	crt  = "../certificates/server.cer"
	key  = "../certificates/server.key"
	ca   = "../certificates/ca.cer"
)

// server is used to implement api.YourService.
type server struct{}

// Echo implements YourService.Echo
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.Value)
	return &pb.StringMessage{Value: "Hello "}, nil
}

// GetSome implements YourService.GetSome
func (s *server) GetSome(ctx context.Context, in *pb.GetMessageRequest) (*pb.StringMessage, error) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Received: %v-%v", in.GetMessageId(), in.GetRevision()))
	if in.GetSub() != nil {
		sb.WriteString(in.GetSub().Subfield)
	}
	log.Printf(sb.String())
	return &pb.StringMessage{Value: sb.String()}, nil
}

func main() {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("could not load server key pair: %s", err)
	}
	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	// Define the connection port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})

	// Greate, define and start the gRPC service
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterYourServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
