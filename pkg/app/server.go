package app

import (
	"github.com/colt005/TrackPixel/pkg/api"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type Server struct {
	router         *gin.Engine
	trackerService api.TrackerService
}

func NewServer(router *gin.Engine, trackerService api.TrackerService) *Server {
	return &Server{
		router:         router,
		trackerService: trackerService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run(":8000")

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
