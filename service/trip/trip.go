package trip

import (
	"context"
	"log"

	"github.com/qiyue365/coolcar/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedTripServiceServer
}

func (s *Service) GetTrip(ctx context.Context, req *pb.GetTripRequest) (*pb.GetTripResponse, error) {
	if ctx.Err() == context.Canceled {
		log.Println("context canceled")
		return nil, status.Error(codes.Canceled, codes.Canceled.String())
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("context deadlineExceeded")
		return nil, status.Error(codes.DeadlineExceeded, codes.DeadlineExceeded.String())
	}
	log.Println("receive request:", req)
	return &pb.GetTripResponse{
		Id: req.Id,
		Trip: &pb.Trip{
			Start:       "abc",
			End:         "xyz",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &pb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &pb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*pb.Location{
				{Latitude: 31, Longitude: 119},
				{Latitude: 33, Longitude: 116},
			},
			Status: pb.TripStatus_IN_PROGRESS},
	}, nil
}
