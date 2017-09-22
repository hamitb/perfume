package main

import (
	"flag"
	"net/http"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	gw "perfumepb"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"golang.org/x/net/context"
)

var (
	echoEndpoint = flag.String("endpoint", "localhost:18870", "endpoint of PerfumeD")
)

func run() error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterPerfumeServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)

	if err != nil {
		return err
	}

	logrus.Infoln("Binding 8081 for Gateway...")

	return http.ListenAndServe(":8081", c.Handler(mux))
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
