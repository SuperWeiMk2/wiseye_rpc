package wiseye_grpc_client

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/SuperWeiMk2/wiseye_rpc" // 替换成实际的Go包路径
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	// 设置DNS resolver，默认情况下gRPC使用plaintext，这里仅作演示
	resolver.SetDefaultScheme("dns")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewWiseyeGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	// 读取本地文件内容
	content, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		log.Fatal(err)
	}

	// 构造并发送FileContent请求
	request := &pb.FileContent{Content: content}
	_, err = c.UploadFile(context.Background(), request)
	if err != nil {
		log.Fatalf("could not upload file: %v", err)
	}
}
