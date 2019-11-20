package daemon

import (
	"log"
	"net/http"
	"os"

	"github.com/dnguy078/go-detector/endpoints"
)

type Server struct {
	router *http.ServeMux
	logger *log.Logger

	host string
	port string
}

func New() (*Server, error) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	router := http.NewServeMux()
	dh := &endpoints.DetectHandler{}

	router.HandleFunc("/detect", dh.Detect)
	return &Server{
		router: router,
		logger: logger,
	}, nil
}

func (s *Server) Start() error {
	s.logger.Println("starting server")
	return http.ListenAndServe(":3000", s.router)
}
