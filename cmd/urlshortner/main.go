package urlshortner

import (
	"fmt"
	"net"
	"os"
	urlshortnerapiv1 "url-shortner/api/grpc/url-shortner/v1"
	"url-shortner/internal/controller"

	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	urlShortnerCtrl := controller.NewURLShortnerController()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	urlshortnerapiv1.RegisterURLShortnerServiceServer(server, urlShortnerCtrl)

	return server.Serve(lis)
}
