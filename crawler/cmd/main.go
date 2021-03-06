package main

import (
	"context"
	"github.com/Elderly-AI/observer/crawler/internal/pkg/vk"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	iCrawler "github.com/Elderly-AI/observer/crawler/internal/app/crawler"
	fCrawler "github.com/Elderly-AI/observer/crawler/internal/pkg/crawler"
	pbCrawler "github.com/Elderly-AI/observer/crawler/pkg/proto/crawler"
)

func registerServices(s *grpc.Server) {
	vkClient, err := vk.New("")
	if err != nil {
		log.Fatalln("Failed to init vk client:", err)
	}
	crawlerFacade := fCrawler.New(vkClient)
	crawlerImplementation := iCrawler.New(crawlerFacade)
	pbCrawler.RegisterCrawlerServer(s, &crawlerImplementation)
}

func newGateway(ctx context.Context, conn *grpc.ClientConn, opts []runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)

	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		//pbAuth.RegisterAuthHandler,
	} {
		err := f(ctx, mux, conn)
		if err != nil {
			return nil, err
		}
	}
	return mux, nil
}

type Options struct {
	Addr string
	Mux  []runtime.ServeMuxOption
}

func createInitialOptions() Options {
	opts := Options{}
	return opts
}

func addGRPCMiddlewares(opts Options) Options {
	return opts
}

func main() {
	//opts := createInitialOptions()
	//opts = addGRPCMiddlewares(opts)

	lis, err := net.Listen("tcp", ":8080") // TODO move to config
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(grpc.MaxSendMsgSize(1024 * 1024 * 1024))
	registerServices(s)

	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatalln(s.Serve(lis))
}
