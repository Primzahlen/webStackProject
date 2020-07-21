package main

import (
	"google.golang.org/grpc"
	"grpc_v1_client/pbfiles"
	"grpc_v1_client/services"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal("Fail to connect grpc server #{err}\n")
	}

	defer conn.Close()
	userClient := pbfiles.NewUserServiceClient(conn)
	services.OneWayLogin(userClient)
	//services.LoginByStream(userClient)
	//userRes, err := userClient.LoginUser(context.Background(), &pbfiles.User{
	//	Email: "ycx@gla.ac.uk",
	//	Name: "Ye Caixu",
	//	Password: "123456",
	//})
	//
	//if err != nil{
	//	log.Fatalf("Failed to request RPC server #{err}\n")
	//}
	//
	//if userRes == nil || userRes.Code != 200 || userRes.User == nil {
	//	log.Fatalf("grpc response is wrong: %v", userRes)
	//}
	//fmt.Println(userRes.User)


	//调用transfer函数
	transferClient := pbfiles.NewTransferServiceClient(conn)
	services.SendMessage(transferClient)

}
