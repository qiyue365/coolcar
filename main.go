package main

import (
	"encoding/json"
	"fmt"

	"github.com/qiyue365/coolcar/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	trip := &pb.Trip{
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
		Status: pb.TripStatus_IN_PROGRESS,
	}
	fmt.Println(trip)

	buf, err := proto.Marshal(trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", buf)
	fmt.Println("len of protobuf: ", len(buf))

	var trip2 pb.Trip
	if err := proto.Unmarshal(buf, &trip2); err != nil {
		panic(err)
	}
	fmt.Println(&trip2)

	buf, err = json.Marshal(&trip2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", buf)
	fmt.Printf("%s\n", buf)
	fmt.Println("len of json: ", len(buf))
}
