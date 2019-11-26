package daemon

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dnguy078/go-detector/adapter"
	"github.com/dnguy078/go-detector/endpoints"
	"github.com/dnguy078/go-detector/service"
)

type Server struct {
	h      *http.Server
	router *http.ServeMux
	host   string
	port   string
	quit   chan bool
	db     service.LoginStorage
	g      service.GeoIPer
}

func New(port string, dbPath string, dbSchemaPath string, geoSeedPath string) (*Server, error) {
	router := http.NewServeMux()
	db, err := adapter.NewDB(dbPath)
	if err != nil {
		return nil, err
	}
	if err := db.LoadFromFile(dbSchemaPath); err != nil {
		return nil, err
	}
	log.Println("database loaded")

	g, err := adapter.NewGeoIP(geoSeedPath)
	if err != nil {
		return nil, err
	}

	detectionSvc := service.NewGeoSuspiciousSvc(db, g)

	dh := &endpoints.DetectHandler{
		DetectSvc: detectionSvc,
	}
	h := &http.Server{Addr: port, Handler: router}

	router.HandleFunc("/detect", endpoints.WithLogging(dh.Detect))
	return &Server{
		router: router,
		port:   port,
		h:      h,
		db:     db,
		g:      g,
		quit:   make(chan bool),
	}, nil
}

func (s *Server) Start() error {
	log.Printf("Starting http service on %s", s.port)
	// Start server
	go func() {
		if err := s.h.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Stop(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	log.Println("gracefully shutting down")
	err := s.db.Close()
	err = s.g.Close()
	err = s.h.Shutdown(ctx)
	return err
}
