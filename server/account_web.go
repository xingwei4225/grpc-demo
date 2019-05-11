package main

import (
	"context"
	"flag"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	gw "tri.com/grpc-demo/model"
	"tri.com/grpc-demo/swagger"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run()error  {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()



	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterAccountServiceHandlerFromEndpoint(ctx,gwmux,*echoEndpoint,opts)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/",gwmux)
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.ServeFile(w,request,"/Users/xingwei/workspace/workspace-go/src/tri.com/grpc-demo/model/account.swagger.json")
	})
	serveSwaggerUI(mux)

	return http.ListenAndServe(":8888",mux)

}

func serveSwaggerUI(mux *http.ServeMux)  {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func main()  {
	flag.Parse()
	if err := run();err != nil{
		log.Fatalln("start server  error:",err)
	}else {
		log.Println("server is starting at 8080...")
	}
}
