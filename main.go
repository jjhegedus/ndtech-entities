package main

import (
	"fmt"
	"io"
	"net/http"
	// entitiesProto "github.com/22ndtech/ndtech-entities/protos/entities"
	// entities "github.com/22ndtech/ndtech-entities/server/entities"
	// "google.golang.org/grpc"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")

	// entityData := entities.EntityData{}
	// fmt.Println("entityData = ", entityData)
}

func main() {
	// grpcServer := grpc.NewServer()
	// fmt.Println("grpcServer:", grpcServer)
	// entitiesServer := entities.NewEntitiesServer()
	// entitiesProto.RegisterEntitiesServer(grpcServer, entitiesServer)

	// l, err := net.Listen("tcp", ":9000")
	// if err != nill{
	// 	fmt.Println("Error Unable to listen on tcp:9000")
	// 	os.Exit(1)
	// }

	// grpcServer.Serve()

	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)

	// type server struct{}
	// type healthServer struct{}
}
