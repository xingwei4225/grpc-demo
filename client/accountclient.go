package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "tri.com/grpc-demo/model"
)

var(
	serverAddr         = flag.String("server_addr", "127.0.0.1:9090", "The server address in the format of host:port")
)

func getAccountInfo(client pb.AccountServiceClient)  {
	ctx,cancel := context.WithTimeout(context.Background(),10 * time.Second)
	defer cancel()
	aid := pb.AccountId{Id:1002}
	ainfo,err:= client.GetAccountInfo(ctx,&aid)
	if err != nil {
		log.Println("call remote service method[GetAccountInfo] error",ainfo)
	}
	fmt.Println(ainfo.Id,ainfo.Age,ainfo.Name,ainfo.Address)
}

func main()  {
	//空的切片，切片的长度和容量都为0，指向底层数组的指针为nil
	var opts []grpc.DialOption
	opts = append(opts,grpc.WithInsecure())

	conn,err := grpc.Dial(*serverAddr,opts...)
	if err != nil {
		log.Fatalln("connect error",err)
	}
	defer conn.Close()
	client := pb.NewAccountServiceClient(conn)
	getAccountInfo(client)
	
}


