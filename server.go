package kompl

import (
	"fmt"
	"net/http"

	"github.com/mitsuse/kompl/predictor"
)

type Server struct {
	port      string
	predictor *predictor.Predictor
}

func NewServer(port string, predictor *predictor.Predictor) *Server {
	// TODO: Configure a server.
	s := &Server{
		port:      port,
		predictor: predictor,
	}

	return s
}

func (s *Server) Port() string {
	return s.port
}

func (s *Server) Predictor() *predictor.Predictor {
	return s.predictor
}

func (s *Server) Run() error {
	http.HandleFunc("/candidates", s.getCandidates)
	address := fmt.Sprintf(":%s", s.Port())

	return http.ListenAndServe(address, nil)
}
