package compl

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	port  string
	model *Model
}

func NewServer(port string, model *Model) *Server {
	// TODO: Configure a server.
	s := &Server{
		port:  port,
		model: model,
	}

	return s
}

func (s *Server) Port() string {
	return s.port
}

func (s *Server) Model() *Model {
	return s.model
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

	var contextSeq []string
	if err := json.Unmarshal([]byte(contextSeqJson), &contextSeq); err != nil {
		// TODO: Return error response.
		return
	}

	candSeq := s.Model().Predict(contextSeq, 10)

	encoder := json.NewEncoder(writer)
	encoder.Encode(candSeq)
}
