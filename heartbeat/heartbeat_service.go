package heartbeat

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"raft/config"
	"raft/heartbeat/rpc/pb"
	"strconv"
	"sync"
	"time"
)

type HeartbeatService struct {
	pb.UnimplementedHeartbeatServiceServer
	client Client
}

func StartRpc() {
	if config.GConfig.Info.Type != "slave" {
		return
	}

	heartbeatServer := &HeartbeatService{
		client: Client{
			lock:            sync.Mutex{},
			nextTimeOutTime: time.Now(),
			Ip:              "127.0.0.1:" + strconv.Itoa(config.GConfig.Heartbeat.Port),
		},
	}

	heartbeatServer.freshTimeOutTime()

	go heartbeatServer.CreateRpc()
	go heartbeatServer.judgeTimeOut()
}

func (s *HeartbeatService) CreateRpc() {
	// 创建启动 rpc 端口
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.GConfig.Heartbeat.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	pb.RegisterHeartbeatServiceServer(serve, s)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *HeartbeatService) Ping(ctx context.Context, in *pb.PingReq) (*pb.PongRes, error) {
	// 接收到心跳之后
	// 刷新 自己的过期时长
	// 并 发送 pong 通知 主节点 存活
	s.freshTimeOutTime()
	log.Println("ping " + strconv.Itoa(config.GConfig.Heartbeat.Port))

	return &pb.PongRes{Ip: strconv.Itoa(config.GConfig.Heartbeat.Port), Message: "pong"}, nil
}

func (s *HeartbeatService) freshTimeOutTime() {
	s.client.lock.Lock()
	defer s.client.lock.Unlock()
	s.client.nextTimeOutTime = time.Now().Add(timeout * time.Second)
}

func (s *HeartbeatService) judgeTimeOut() {
	for {
		time.Sleep(500 * time.Millisecond)
		log.Printf("%v judge is TimeOut ? \n", s.client.nextTimeOutTime)

		if time.Now().After(s.client.nextTimeOutTime) {
			log.Printf("%v is TimeOut . \n", s.client.nextTimeOutTime)
			return
		}
	}
}

func StartPingService() {
	if config.GConfig.Info.Type != "master" {
		return
	}

	for {
		for index, ipServe := range GIpServe {
			//if !ipServe.IsAlive {
			//	// 该节点已下线
			//	continue
			//}

			// 向 从节点 发送心跳
			conn, err := grpc.Dial(ipServe.Ip, grpc.WithInsecure())
			if err != nil {
				log.Print("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewHeartbeatServiceClient(conn)
			res, err := c.Ping(context.Background(), &pb.PingReq{Ip: ipServe.Ip})
			if err != nil {
				GIpServe[index].IsAlive = false
				log.Print("did not connect: %v", err)
			}
			log.Println(res.GetMessage())
		}

		time.Sleep(intervalTime * time.Second)
	}
}
