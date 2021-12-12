package pkg

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/barklan/fan-in-fan-out/protos"
	"google.golang.org/grpc"
)

var port = 50051

type server struct {
	pb.UnimplementedReporterServer
}

func (s *server) Report(ctx context.Context, in *pb.ReportRequest) (*pb.ReportReply, error) {
	reportMessage := in.GetMessage()
	log.Printf("Received: %v", reportMessage)
	_ = in.GetToken()

	return &pb.ReportReply{Message: "ok"}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReporterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Panicf("failed to serve: %v", err)
	}
}
