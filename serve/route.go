package serve

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"raft/heartbeat"
	"raft/heartbeat/rpc/pb"
	"strconv"
)

func route(r *gin.Engine) {
	r.GET("/op", func(c *gin.Context) {
		op := c.Query("op")
		number, _ := strconv.Atoi(c.Query("number"))
		failure := 0

		for _, ipServe := range heartbeat.GIpServe {
			conn, err := grpc.Dial(ipServe.Ip, grpc.WithInsecure())
			if err != nil {
				log.Print("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewHeartbeatServiceClient(conn)
			res, err := c.SendOp(context.Background(), &pb.OpReq{Op: op, Number: int32(number)})
			if !res.Flag {
				failure += 1
			}
		}

		if failure > len(heartbeat.GIpServe)/2 {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})
}
