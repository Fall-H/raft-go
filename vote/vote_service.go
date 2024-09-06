package vote

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"raft/config"
	"raft/vote/rpc/pb"
	"strconv"
)

type VoteService struct {
	pb.UnimplementedVoteServer
	vote int32
}

func (s *VoteService) SendVote(ctx context.Context, in *pb.VoteReq) (*pb.VoteRes, error) {
	// 接收到 vote 命令
	// 清除 从节点信息 进入 投票模式
	// 判断 任期信息

	term := in.Term
	name := in.Name
	vote := in.Vote

	if config.GConfig.Info.Term >= in.Term {
		term = config.GConfig.Info.Term
		name = config.GConfig.Info.Name
	}

	if s.vote > vote {
		vote = s.vote
	}

	return &pb.VoteRes{Name: name, Term: term, Vote: vote}, nil
}

func CreateVote() {
	if config.GConfig.Info.Type != "vote" {
		return
	}

	// 向其他节点发送 请求他们向 自己投票
}

func (s *VoteService) CreateVoteServe() {
	// 一直启动 服务端
	// 如果出现主节点挂了 从节点有可能发送 vote 选票包
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(config.GConfig.Vote.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	pb.RegisterVoteServer(serve, s)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func CreateVoteClient() {
	config.GConfig.Info.Term += 1
	vote := int32(1)

	for index, ipServe := range GIpServe {
		conn, err := grpc.Dial(ipServe.Ip, grpc.WithInsecure())
		if err != nil {
			log.Print("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewVoteClient(conn)
		res, err := c.SendVote(context.Background(), &pb.VoteReq{Name: config.GConfig.Info.Name,
			Term: config.GConfig.Info.Term, Vote: vote})
		if err != nil {
			GIpServe[index].IsAlive = false
			log.Print("did not connect: %v", err)
		}

		// 判断选票
	}
}
