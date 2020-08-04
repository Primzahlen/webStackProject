package services

import (
	"bytes"
	"encoding/json"
	"http_client/model"
	"log"
	"net/http"
	"reflect"
)

func SendMessage(client *http.Client) {
	message := initTransferMessage()
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(message)
	//m_size, _ := proto.Marshal(message)
	//fmt.Printf("buf size: %d bytes\n", buf.Len())
	response, err :=client.Post("http://127.0.0.1:60001/DataTransmission", "application/json", buf)
	if err != nil {
		log.Fatalf("Http request failed: #{err}\n")
	}
	defer response.Body.Close()

	// 格式转换
	var target model.DataResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("无法解析JSON: #{err}\n")
	}
	if target.Code != 200 || target.Message == "" {
		log.Fatalf("接收数据错误！: #{err}\n")
	}
	//fmt.Println(target.Data.Field1)
}

func initTransferMessage() *model.BenchmarkMessage {
	// 设置MESSAGE中有关类型的赋值
	b := false
	var i32 int32 = 10000
	var i64 int64 = 10000
	var s = "I am a student from University of Glasgow, I want to be a good programmer."
	var message model.BenchmarkMessage
	// 使用反射，V 获取message所有元素
	v := reflect.ValueOf(&message).Elem()
	// num：message中field数量
	num := v.NumField()
	//遍历message的每一个field
	for i:=0; i<num; i++ {
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