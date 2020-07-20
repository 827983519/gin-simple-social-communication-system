package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	DBagent "sbs-entrytask-template/agent/db"
	grpc_follow "sbs-entrytask-template/apps/example/grpc/proto/follow_proto"
	_ "sbs-entrytask-template/apps/example/grpc/proto/follow_proto"
	"sbs-entrytask-template/apps/example/grpc/proto/user_proto"
	"sbs-entrytask-template/apps/example/grpc/server"
)



func main() {
	var err error
	DBagent.DB, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_entry_task?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	defer DBagent.DB.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		print("what'")
	}
	s := grpc.NewServer()
	grpc_user.RegisterUserServiceServer(s, &server.UserServer{})
	grpc_follow.RegisterFollowServiceServer(s, &server.FollowServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}