package vote

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"raft/config"
	"raft/model"
	"raft/observer"
	"raft/op"
	"raft/vote/rpc/pb"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	pb.UnimplementedVoteServer
	observer.Observer
	vote          int32
	state         string
	ObserverState observer.State
}

func (s *Service) SendVote(ctx context.Context, in *pb.VoteReq) (*pb.VoteRes, error) {
	// 接收到 vote 命令
	// 清除 从节点信息 进入 投票模式
	// 判断 任期信息

	term := in.Term
	name := in.Name
	vote := int32(0)

	// 如何自己的 任期 大于对方任期 则对方直接服从你的任期
	if config.GConfig.Info.Term >= in.Term {
		term = config.GConfig.Info.Term
		name = config.GConfig.Info.Name
	}

	if s.vote > vote {
		vote = s.vote
	}

	return &pb.VoteRes{Name: name, Term: term, Vote: vote, State: s.state, Ip: config.GConfig.Heartbeat.Port}, nil
}

func (s *Service) ChangeStateSlave(ctx context.Context, in *pb.VoteReq) (*pb.VoteRes, error) {
	s.state = model.StateSlave
	config.GConfig.Heartbeat.MasterIp = in.Ip
	config.GConfig.Info.Term = in.Term
	s.ObserverState.NotifyAll(s.state)

	go syncOpList(in.OpList)

	return &pb.VoteRes{}, nil
}

func syncOpList(opList []string) {
	model.Number = 0

	for _, item := range opList {
		temp := strings.Split(item, ":")
		number, _ := strconv.Atoi(temp[1])
		op.Op(temp[0], int32(number))
	}

	model.OpList = opList
}

func (s *Service) CreateVoteServe() {
	// 一直启动 服务端
	// 如果出现主节点挂了 从节点有可能发送 vote 选票包
	lis, err := net.Listen("tcp", config.GConfig.Vote.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serve := grpc.NewServer()
	pb.RegisterVoteServer(serve, s)
	if err := serve.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Service) Update(state string) {
	s.state = state

	go s.CreateVote()
}

func (s *Service) GetState() string {
	return s.state
}

func (s *Service) CreateVote() {
	if s.state != model.StateVote {
		return
	}

	// 向其他节点发送 请求他们向 自己投票
	s.CreateVoteClient()
}

func (s *Service) CreateVoteClient() {
	config.GConfig.Info.Term += 1

	for {
		s.vote = int32(1)
		for index, ipServe := range GIpServe {
			res := makeVoteReq(index, ipServe, s.vote)
			if res == nil {
				continue
			}

			// 判断任期 如果大于自己任期
			if res.Term > config.GConfig.Info.Term && res.State == model.StateMaster {
				config.GConfig.Heartbeat.MasterIp = res.Ip
				s.ObserverState.NotifyAll(model.StateSlave)
				return
			}

			// 获得相应 选票加一
			s.vote += 1
			s.state = model.StateVote

			if s.judgeVote(s.vote, res.Vote, ipServe) {
				return
			} else {
				break
			}
		}
	}
}

func makeVoteReq(index int, ipServe model.IpServe, vote int32) *pb.VoteRes {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial(ipServe.Ip, grpc.WithInsecure())
	if err != nil {
		log.Print("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewVoteClient(conn)
	res, err := c.SendVote(ctx, &pb.VoteReq{Name: config.GConfig.Info.Name,
		Term: config.GConfig.Info.Term, Vote: vote})
	if err != nil {
		GIpServe[index].IsAlive = false
		log.Print("did not connect: %v", err)
	}

	return res
}

func (s *Service) judgeVote(owenVote int32, otherVote int32, ipServe model.IpServe) bool {
	// 判断选票
	if otherVote == owenVote {
		// 选票一致 重新投票
		time.Sleep(model.Timeout)
		return false
	} else if otherVote < owenVote {
		conn, err := grpc.Dial(ipServe.Ip, grpc.WithInsecure())
		if err != nil {
			log.Print("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewVoteClient(conn)

		// 如果自己选票 大于其他人的选票 则自己成为 master
		config.GConfig.Info.Term += 1
		c.ChangeStateSlave(context.Background(), &pb.VoteReq{Term: config.GConfig.Info.Term, Ip: config.GConfig.Heartbeat.Port})
		s.ObserverState.NotifyAll(model.StateMaster)
		return true
	} else if otherVote > owenVote {
		// 如果自己选票 小于其他人的选票 则自己成为 被投过 票者
		s.ObserverState.NotifyAll(model.StateVoted)
		return true
	}
	return false
}
