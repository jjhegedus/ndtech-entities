package entities

import (
	"fmt"
	// "io"
	// "net/http"

	protos "github.com/jjhegedus/ndtech-entities/protos"
	// "google.golang.org/grpc"
	"context"
)

type EntitiesServer struct {
}

func NewEntitiesServer() *EntitiesServer {
	return &EntitiesServer{}
}

func (es *EntitiesServer) CreateEntity(ctx context.Context, ed *protos.EntityData) (*protos.Entity, error) {
	fmt.Println("EntitiesServer: CreateEntities: EntityData: ", ed)
	return &protos.Entity{EntityData: ed}, nil
}

func (es *EntitiesServer) GetEntity(ctx context.Context, id *protos.EntityId) (*protos.Entity, error) {
	fmt.Println("EntitiesServer: GetEntity: EntityId: ", id)
	return &protos.Entity{}, nil
}
