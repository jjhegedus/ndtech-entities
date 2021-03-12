package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	entities "github.com/jjhegedus/ndtech-entities/entities"
	protos "github.com/jjhegedus/ndtech-entities/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")

	entityData := protos.EntityData{}
	fmt.Println("entityData = ", entityData)
}

func main() {
	grpcServer := grpc.NewServer()
	fmt.Println("grpcServer:", grpcServer)
	entitiesServer := entities.NewEntitiesServer()
	protos.RegisterEntitiesServer(grpcServer, entitiesServer)

	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error Unable to listen on tcp:9000")
		os.Exit(1)
	}

	grpcServer.Serve(l)

	// portNumber := "9000"
	// http.HandleFunc("/", handle)
	// fmt.Println("Server listening on port ", portNumber)
	// http.ListenAndServe(":"+portNumber, nil)

	// type server struct{}
	// type healthServer struct{}
}
