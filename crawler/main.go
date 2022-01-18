package main

import (
	"bytes"
	"context"
	"fmt"
	pbCrawler "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
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
			Users: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75},
		},
	)

	if err != nil {
		log.Fatalf("could not get answer: %v", err)
	}

	for _, photo := range resp.Photos {
		for j, p := range photo.Photos {
			file, _ := os.Create(fmt.Sprintf("%d_%d.jpg", photo.User, j))
			io.Copy(file, bytes.NewReader(p))
			file.Close()
		}
	}
}
