package heartbeat

import (
	"context"
	"raft/heartbeat/rpc/pb"
)

type HeartbeatService struct {
	pb.UnimplementedHeartbeatServer
}

func (h *HeartbeatService) PingRpc(c context.Context, req *pb.PingRpcModel) (*pb.PongRpcModel, error) {

	return &pb.PongRpcModel{}, nil
}
