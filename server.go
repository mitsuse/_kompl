package compl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	port      string
	predictor *Predictor
}

func NewServer(port string, predictor *Predictor) *Server {
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

func (s *Server) Predictor() *Predictor {
	return s.predictor
}

func (s *Server) Run() error {
	http.HandleFunc("/", s.handler)
	address := fmt.Sprintf(":%s", s.Port())

	return http.ListenAndServe(address, nil)
}

func (s *Server) handler(writer http.ResponseWriter, requst *http.Request) {
	header := writer.Header()
	header.Set("Content-Type", "application/json")

	contextSeqJson := requst.FormValue("context")
	prefix := requst.FormValue("prefix")

	var contextSeq []string
	if err := json.Unmarshal([]byte(contextSeqJson), &contextSeq); err != nil {
		// TODO: Return error response.
		return
	}

	candSeq := s.Predictor().Predict(contextSeq, prefix, 10)

	encoder := json.NewEncoder(writer)
	encoder.Encode(candSeq)
}
