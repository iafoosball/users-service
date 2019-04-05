package main

import (
	"context"
	"fmt"
	"github.com/iafoosball/users-service/usersgrpc/server/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"

	pb "github.com/iafoosball/users-service/usersgrpc/userspb"
)

func main() {
	//flag.Parse() // parse our flags

	// initialize grpc service
	serv := &UsersServiceServer{}

	//port := getEnv("SERVICE_ADDR", "localhost:8001")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, serv)

	// determine whether to use TLS
	log.Printf("%T is listening...", serv)
	grpcServer.Serve(lis)
}

// getEnv returns environmental variable called name, or fallback if empty
func getEnv(name string, fallback string) string {
	v, ok := os.LookupEnv(name)
	if ok {
		return v
	}
	return fallback
}

// UsersService Read Model
type UsersServiceServer struct {}

func (*UsersServiceServer) ReadOne(ctx context.Context, req *pb.ReadOneRequest) (*pb.ReadOneResponse, error) {
	data, err := redis.HGETALL(req.Userhash)
	if err != nil {
		log.Fatal("Failed fetching redis data: "+err.Error())
		return &pb.ReadOneResponse{}, err
	}

	// create goroutine that polls matches on streaming endpoint and gets all the changes? matches should calculate?
	dmap := data.(map[string]string)
	u := &pb.User{
		Userhash: req.Userhash,
		Nickname: dmap["nickname"],
		Mmr: sti32(dmap["mmr"]),
		Winrate: sti32(dmap["winrate"]),
		Totalgoals: sti32(dmap["totalgoals"]),
	}
	return &pb.ReadOneResponse{User:u}, nil
}

func sti32(s string) int32 {
	mmr64, err := strconv.ParseInt(s,10,32)
	if err != nil {
		panic(err)
	}
	return int32(mmr64)
}

