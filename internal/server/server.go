package server

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	s.initLogger()
}
