package h_grpc

import (
	convert_proto "github.com/fishmanDK/internal/convert-proto/proto"
	"github.com/fishmanDK/internal/service"
	"google.golang.org/grpc"
)

type ServerApi struct {
	convert_proto.UnimplementedAuthServer
	Service *service.Service
}

func Register(gRPC *grpc.Server, service *service.Service) {
	convert_proto.RegisterAuthServer(gRPC, &ServerApi{Service: service})
}

func NewServer() *ServerApi {
	return &ServerApi{}
}
