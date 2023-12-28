package server

import (
	"fmt"

	"net/http"

	slog "github.com/eric-tech01/simple-log"

	"github.com/eric-tech01/taurus/pkg/conf"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	config *Config
}

func New() *Server {
	config := &Config{Host: conf.GetString("taurus_server_http.Host"), Port: conf.GetInt("taurus_server_http.Port")}
	s := &Server{Engine: gin.New(), config: config}
	return s
}

func (s *Server) Start() error {
	for _, route := range s.Engine.Routes() {
		slog.Info("add route ", route.Method, route.Path)
	}
	err := s.Run(s.config.Address())
	if err == http.ErrServerClosed {
		slog.Panic("close ", s.config.Address(), err)
		return nil
	}
	return err
}

type Config struct {
	Host string
	Port int
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
