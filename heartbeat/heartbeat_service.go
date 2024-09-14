package heartbeat

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"raft/config"
	"raft/heartbeat/rpc/pb"
	"raft/model"
	"raft/observer"
	"raft/op"
	"runtime"
	"strconv"
	"time"
)

type Service struct {
	pb.UnimplementedHeartbeatServiceServer
	observer.Observer
	Client        model.Client
	state         string
	ObserverState observer.State
}

func (s *Service) startPingService() {
	if s.state != model.StateMaster {
		return
	}

	for {
		for index, ipServe := range GIpServe {
			//if !ipServe.IsAlive {
			//	// 该节点已下线
			//	continue
			//}
			clientDeadline := time.Now().Add(3 * time.Second)
			ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
			defer cancel()

			// 向 从节点 发送心跳
			conn, err := grpc.DialContext(ctx, ipServe.Ip, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Print("did not connect: %v", err)
			}
			defer conn.Close()

			c := pb.NewHeartbeatServiceClient(conn)
			res, err := c.Ping(ctx, &pb.PingReq{Ip: ipServe.Ip})
			if err != nil {
				GIpServe[index].IsAlive = false
				log.Print("did not connect: %v", err)
			} else {
				log.Println(res.GetMessage())
			}
		}

		log.Println(runtime.NumGoroutine())

		time.Sleep(model.IntervalTime)
	}
}

func (s *Service) StartRpc() {
	if s.state != model.StateSlave {
		return
	}

	s.freshTimeOutTime()

	go s.judgeTimeOut()
}

func (s *Service) CreateRpc() {
	// 创建启动 rpc 端口
	lis, err := net.Listen("tcp", config.GConfig.Heartbeat.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	pb.RegisterHeartbeatServiceServer(serve, s)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Service) Ping(ctx context.Context, in *pb.PingReq) (*pb.PongRes, error) {
	// 接收到心跳之后
	// 刷新 自己的过期时长
	// 并 发送 pong 通知 主节点 存活
	s.freshTimeOutTime()
	log.Println("ping " + config.GConfig.Heartbeat.Port)

	return &pb.PongRes{Ip: config.GConfig.Heartbeat.Port, Message: "pong"}, nil
}

func (s *Service) SendOp(ctx context.Context, in *pb.OpReq) (*pb.OpRes, error) {
	log.Println(" " + in.Op + ":" + strconv.Itoa(int(in.Number)))
	s.freshTimeOutTime()

	op.Op(in.Op, in.Number)
	model.OpList = append(model.OpList, in.Op+":"+strconv.Itoa(int(in.Number)))

	return &pb.OpRes{Flag: true}, nil
}

func (s *Service) freshTimeOutTime() {
	s.Client.Lock.Lock()
	defer s.Client.Lock.Unlock()
	s.Client.NextTimeOutTime = time.Now().Add(model.Timeout)
}

func (s *Service) judgeTimeOut() {
	for {
		time.Sleep(model.PollingTime)
		log.Printf("%v judge is TimeOut ? \n", s.Client.NextTimeOutTime)

		// 主节点超时 变成投票模式
		if time.Now().After(s.Client.NextTimeOutTime) {
			log.Printf("%v is TimeOut . \n", s.Client.NextTimeOutTime)
			s.state = model.StateVote
			s.ObserverState.NotifyAll(s.state)
			return
		}
	}
}

func (s *Service) Update(state string) {
	s.state = state

	go s.startPingService()
	go s.StartRpc()
}

func (s *Service) GetState() string {
	return s.state
}
