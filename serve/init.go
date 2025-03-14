package serve

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"raft/config"
	"raft/model"
	"raft/observer"
)

type Service struct {
	state         string
	ObserverState observer.State
}

func (s *Service) CreateServe() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		if s.state == model.StateSlave {
			c.Redirect(http.StatusMovedPermanently, config.GConfig.Heartbeat.MasterIp)
		}
	})

	route(r)

	r.Run(config.GConfig.Serve.Ip)
}

func (s *Service) Update(state string) {
	s.state = state
}

func (s *Service) GetState() string {
	return s.state
}
