package model

// Response type
type Response struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Data *BenchmarkMessage `json:"data"`
}
// fixed size data
type BenchmarkMessage struct {
	Field1 string `json:"Field1"`
	Field9 string `json:"Field9"`
	Field18 string `json:"Field18"`
	Field80 bool `json:"Field80"`
	Field81 bool `json:"Field81"`
	Field2 int32 `json:"Field2"`
	Field3 int32 `json:"Field3"`
	Field280 int32 `json:"Field280"`
	Field6 int32 `json:"Field6"`
	Field22 uint64 `json:"Field22"`
	Field4 string `json:"Field4"`
	Field5 []uint64 `json:"Field5"`
	Field59 bool `json:"Field59"`
	Field7 string `json:"Field7"`
	Field16 int32 `json:"Field16"`
	Field130 int32 `json:"Field130"`
	Field12 bool `json:"Field12"`
	Field17 bool `json:"Field17"`
	Field13 bool `json:"Field13"`
	Field14 bool `json:"Field14"`
	Field104 int32 `json:"Field104"`
	Field100 int32 `json:"Field100"`
	Field101 int32 `json:"Field101"`
	Field102 string  `json:"Field102"`
	Field103 string `json:"Field103"`
	Field29 int32 `json:"Field29"`
	Field30 bool `json:"Field30"`
	Field60 int32 `json:"Field60"`
	Field271 int32 `json:"Field271"`
	Field272 int32 `json:"Field272"`
	Field150 int32 `json:"Field150"`
	Field23 int32 `json:"Field23"`
	Field24 bool `json:"Field24"`
	Field25 int32 `json:"Field25"`
	Field78 bool `json:"Field78"`
	Field67 int32 `json:"Field67"`
	Field68 int32 `json:"Field68"`
	Field128 int32 `json:"Field128"`
	Field129 string `json:"Field129"`
	Field131 int32 `json:"Field131"`
}
