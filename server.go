package compl

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
	// TODO: Run a compl server.
	return nil
}
