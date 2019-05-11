package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "tri.com/grpc-demo/model"
)

type accountServiceServer struct {
	accounts  []*pb.AccountInfo
}

func (s *accountServiceServer) GetAccountInfo(ctx context.Context, req *pb.AccountId) (*pb.AccountInfo, error) {
	log.Println(req.Id)
	for _, accountInfo := range s.accounts {
		if accountInfo.Id == req.Id{
			return accountInfo,nil
		}
	}
	return &pb.AccountInfo{Id:req.Id}, nil
}
func (*accountServiceServer) ListAccountInfos(req *pb.ReqMsg, srv pb.AccountService_ListAccountInfosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAccountInfos not implemented")
}
func (*accountServiceServer) AddAccountInfos(srv pb.AccountService_AddAccountInfosServer) error {
	return status.Errorf(codes.Unimplemented, "method AddAccountInfos not implemented")
}

func (s *accountServiceServer)loadData()  {
	err := json.Unmarshal(exampleData,&s.accounts)
	if err != nil {
		log.Fatalln("load data occur an error:",err)
	}else {
		log.Printf("成功导入了:[%d]条账户信息\n",len(s.accounts))
	}
}

func newServer()*accountServiceServer  {
	s := &accountServiceServer{

	}
	s.loadData()
	return s
}

var(
	port = flag.Int("port",9090,"请设置端口")
)

func main()  {
	flag.Parse()
	lis,err := net.Listen("tcp",fmt.Sprintf("localhost:%d",*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv,newServer())
	srv.Serve(lis)
}

var exampleData = []byte(`[{
		"id":1001,
		"name":"xingwei",
		"age":35,
		"address":"科学大道上铁银欣花园802"
	},{
		"id":1002,
		"name":"wanghailin",
		"age":34,
		"address":"望江西路555号"
	}]`)
