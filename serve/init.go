package serve

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"raft/config"
	"raft/model"
	"raft/observer"
	"time"
)

type Service struct {
	state         string
	ObserverState observer.State
}

func (s *Service) CreateServe() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		if s.state == model.StateSlave {
			targetURL := &url.URL{
				Scheme:   "http",
				Host:     config.GConfig.MasterServe.Ip, // 确保包含端口号，如 "192.168.0.1:8080"
				Path:     c.Request.URL.Path,
				RawQuery: c.Request.URL.RawQuery, // 追加查询参数
			}

			newUrl := targetURL.String()

			// 创建新请求
			req, err := http.NewRequest(c.Request.Method, newUrl, c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
				return
			}
			req.Header = c.Request.Header.Clone()

			// 发送请求到 Master
			client := &http.Client{Timeout: 5 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"error": "Failed to reach Master"})
				return
			}
			defer resp.Body.Close()

			// 将 Master 的响应头复制到客户端
			for key, values := range resp.Header {
				for _, value := range values {
					c.Writer.Header().Add(key, value)
				}
			}

			// 设置状态码并写入响应体
			c.Status(resp.StatusCode)
			body, _ := io.ReadAll(resp.Body)
			c.Writer.Write(body)

			// 中止后续中间件
			c.Abort()
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
