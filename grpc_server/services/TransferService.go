package services

import (
	"context"
	"grpc_v1/pb"
)

type Transfer struct {

}

func (*Transfer) DataTransmission(ctx context.Context, in *pb.BenchmarkMessage) (*pb.BenchmarkMessage, error) {
	// update
	//fmt.Printf("The size of the received packet is %d bytes \n", unsafe.Sizeof(in))
	status := "OK"
	var i int32 = 100
	in.Field1 = status
	in.Field2 = i
	return in, nil
}

func (t *Transfer) DataTransmissionStream(message *pb.BenchmarkMessage, server pb.TransferService_DataTransmissionStreamServer) error {
	panic("implement me")
}