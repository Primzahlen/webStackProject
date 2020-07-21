package services

import (
	"context"
	"fmt"
	"grpc_v1/pbfiles"
	"unsafe"
)

//func RecvMessage(recvServer pbfiles.TransferServiceServer) {
//	status := "OK"
//	var i int32 = 100
//	message := pbfiles.TransferServiceServer()
//}
type Transfer struct {

}

func (*Transfer) DataTransmission(ctx context.Context, in *pbfiles.BenchmarkMessage) (*pbfiles.BenchmarkMessage, error) {
	// update
	//b := make([]byte, 0)
	fmt.Printf("接收数据包的大小为 %d bytes \n", unsafe.Sizeof(in))
	status := "OK"
	var i int32 = 100
	in.Field1 = status
	in.Field2 = i
	return in, nil
}

func (t *Transfer) DataTransmissionStream(message *pbfiles.BenchmarkMessage, server pbfiles.TransferService_DataTransmissionStreamServer) error {
	panic("implement me")
}