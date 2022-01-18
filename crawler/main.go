package main

import (
	"context"
	pbCrawler "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
	"google.golang.org/grpc"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*1024)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pbCrawler.NewCrawlerClient(conn)
	resp, err := client.GetVkUsersPhotosHandler(context.Background(),
		&pbCrawler.GetVkUsersPhotosHandlerRequest{
			Users: []uint64{1},
		},
	)

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	//for _, photo := range resp.Photos {
	//	for j, p := range photo.Photos {
	//		file, _ := os.Create(fmt.Sprintf("%d_%d.jpg", photo.User, j))
	//		io.Copy(file, bytes.NewReader(p))
	//		file.Close()
	//	}
	//}
	resp = resp
}
