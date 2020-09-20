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
	//fmt.Printf("buf size: %d bytes\n", buf.Len())
	response, err :=client.Post("http://127.0.0.1:60001/DataTransmission", "application/json", buf)
	if err != nil {
		log.Fatalf("Http request failed: #{err}\n")
	}
	defer response.Body.Close()
	// Format conversion
	var target model.DataResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&target)
	if decodeErr != nil {
		log.Fatalf("Unable to parse JSON: %v\n", decodeErr)
	}
	if target.Code != 200 || target.Message == "" {
		log.Fatalf("Receiving data error ！: %v \n", target)
	}
}

func initTransferMessage() *model.BenchmarkMessage {
	// Sets the assignment of the type in MESSAGE
	b := false
	var i32 int32 = 10000
	var i64 int64 = 10000
	var s = "I am a student from University of Glasgow, I want to be a good programmer."
	var message model.BenchmarkMessage
	v := reflect.ValueOf(&message).Elem()  // Using reflection, V gets all the elements of the message
	num := v.NumField() // num：field numbers in message
	for i:=0; i<num; i++ {	// Traverse each field of message
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