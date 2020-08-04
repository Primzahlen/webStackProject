package services

import (
	"context"
	"grpc_v1_client/pbfiles"
	"log"
	"reflect"
)

func SendMessage(sendClient pbfiles.TransferServiceClient) {
	message := initTransferMessage()
	//m_size, _ := proto.Marshal(message)
	//fmt.Printf("message size: %d bytes\n\n\n", len(m_size))
	response, err :=sendClient.DataTransmission(context.Background(), message)
	if err != nil || response.Field1 != "OK"{
		log.Fatalf("接收message发生错误 %v", err)
	}
	//fmt.Println(response)
}

func initTransferMessage() *pbfiles.BenchmarkMessage {
	// 设置MESSAGE中有关类型的赋值
	b := false
	var i32 int32 = 10000
	var i64 int64 = 10000
	var s = "I am a student from University of Glasgow, I want to be a good programmer."
	var message pbfiles.BenchmarkMessage
	// 使用反射，V 获取message所有元素
	v := reflect.ValueOf(&message).Elem()
	// num：message中field数量
	num := v.NumField()
	//遍历message的每一个field
	for i:=3; i<num; i++ {
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