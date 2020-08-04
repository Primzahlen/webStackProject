package services

import (
	"encoding/json"
	"http_v1_server/model"
	"log"
	"net/http"
)

type Transfer struct {

}

func DataTransmission(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var message model.BenchmarkMessage
	err := decoder.Decode(&message)
	if err != nil {
		log.Println("Parse packet error！！")
	}

	//fmt.Printf("接收数据包的大小为 %d bytes \n", r.ContentLength)
	defer r.Body.Close()
	w.Header().Set("Conent-type", "application/json")
	// update
	//b := make([]byte, 0)
	status := "OK"
	var i int32 = 100
	message.Field1 = status
	message.Field2 = i
	ResErr := json.NewEncoder(w).Encode(model.Response{
		Code: 200,
		Message: "OK",
		Data: &message,
	})
	if ResErr != nil {
		log.Println("An error occurred encapsulating data！！")
	}
}

