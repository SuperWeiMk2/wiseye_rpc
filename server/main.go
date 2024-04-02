package wiseye_grpc_server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"net"

	pb "github.com/SuperWeiMk2/wiseye_rpc" // 替换成实际的Go包路径
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedWiseyeGrpcServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) UploadFile(ctx context.Context, file *pb.FileContent) (*empty.Empty, error) {
	fmt.Println(string(file.Content)) // 打印接收到的文件内容
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWiseyeGrpcServer(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
