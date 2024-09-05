package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"raft/config"
	"raft/heartbeat"
	"raft/heartbeat/rpc/pb"
	"strconv"
)

func init() {
	go initHeartbeat()
}

func initHeartbeat() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.GConfig.Heartbeat.Port))

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	log.Print("注册心跳服务")
	pb.RegisterHeartbeatServer(srv, &heartbeat.HeartbeatService{})
	log.Print("启动心跳服务")
	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func main() {
}
