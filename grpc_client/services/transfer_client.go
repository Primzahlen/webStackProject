package services

import (
	"context"
	"grpc_v1_client/pb"
	"log"
	"reflect"
)

func SendMessage(sendClient pb.TransferServiceClient) {
	message := initTransferMessage()
	//fmt.Printf("message size: %d bytes\n\n\n", len(m_size))
	response, err :=sendClient.DataTransmission(context.Background(), message)
	if err != nil || response.Field1 != "OK"{
		log.Fatalf("Error receiving message %v", err)
	}
}

func initTransferMessage() *pb.BenchmarkMessage {
	// Sets the assignment of the type in MESSAGE
	b := false
	var i32 int32 = 10000
	var i64 int64 = 10000
	var s = "I am a student from University of Glasgow, I want to be a good programmer."
	var message pb.BenchmarkMessage
	v := reflect.ValueOf(&message).Elem()	// Using reflection, V gets all the elements of the message
	num := v.NumField()	// num：field numbers in message
	for i:=3; i<num; i++ {	// Traverse each field of message
		field := v.Field(i)
		if field.Type().Kind() == reflect.Ptr {
			switch v.Field(i).Type().Elem().Kind() {
			case reflect.Int, reflect.Int32:
				field.Set(reflect.ValueOf(&i32))
			case reflect.Int64:
				field.Set(reflect.ValueOf(&i64))
			case reflect.Bool:
				field.Set(reflect.ValueOf(&b))
			case reflect.String:
				field.Set(reflect.ValueOf(&s))
			}
		}else {
			switch field.Kind() {
			case reflect.Int, reflect.Int32, reflect.Int64:
				field.SetInt(1000000)
			case reflect.Bool:
				field.SetBool(b)
			case reflect.String:
				field.SetString(s)
			}
		}
	}
	return &message
}