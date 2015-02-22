/*
Package "server" provides the K-best word completion server.

This server suggests the completion candidates sorted by N-gram frequency.
*/
package server

import (
	"fmt"
	"net/http"

	"github.com/mitsuse/kompl/predictor"
)

type Server struct {
	port      string
	predictor *predictor.Predictor
}

/*
Create an instance of word completion server.
This requires the port number used to communicate with clients
and the K-best word predictor.
*/
func New(port string, predictor *predictor.Predictor) *Server {
	// TODO: Configure a server.
	s := &Server{
		port:      port,
		predictor: predictor,
	}

	return s
}

/*
Return The port number used to communicate with clients.
*/
func (s *Server) Port() string {
	return s.port
}

/*
Return the the K-best word predictor.
*/
func (s *Server) Predictor() *predictor.Predictor {
	return s.predictor
}

/*
Run the word completion server.
This executes rooting and listen the specified port.
*/
func (s *Server) Run() error {
	http.HandleFunc("/candidates", s.getCandidates)
	http.HandleFunc("/description", s.getDescription)

	address := fmt.Sprintf(":%s", s.Port())

	return http.ListenAndServe(address, nil)
}
