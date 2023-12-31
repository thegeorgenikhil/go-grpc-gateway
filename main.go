package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	helloworldpb "github.com/thegeorgenikhil/go-grpc-gateway/proto/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{
		Message: in.Name + " world",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	s := grpc.NewServer()

	helloworldpb.RegisterGreeterServer(s, &server{})

	log.Println("Serving gRPC on 0.0.0.0:8080")

	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(context.Background(), "0.0.0.0:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial server: ", err)
	}

	gwmux := runtime.NewServeMux()

	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
