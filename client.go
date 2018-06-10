package main

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"fmt"

	pb "github.com/DesmondANIMUS/greeter/greet"
	"google.golang.org/grpc"
)

const address = "localhost:8080"

func main() {
	conn, e := grpc.Dial(address, grpc.WithInsecure())
	err(e)
	defer conn.Close()

	random := randgen(20)

	client := pb.NewPingClient(conn)
	r, e := client.SayHello(context.Background(), &pb.HelloRequest{Name: random})
	err(e)

	fmt.Println(r.Reply)
}

func err(err error) {
	if err != nil {
		panic(err)
	}
}

func randgen(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
