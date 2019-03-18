package usersgrpc

import (
	fmt "fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {
	//flag.Parse() // parse our flags

	// initialize arango DB with cols
	db := &ArangoDB{
		Host: "arangodb",
		Port: 8529,
		User: "root",
		Password: "users-password",
	}

	db.InitDatabase()
	// initialize grpc service
	serv := &usersServiceServer{
		usersCol: db.col(usersColName),
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", db.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterUsersServiceServer(grpcServer, serv)

	// determine whether to use TLS
	grpcServer.Serve(lis)
}
