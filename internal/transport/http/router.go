package http

func (s *Server) Router() {
	u := s.HTTP.Group("/user/")
	u.POST("/add/", s.handler.CreateUser)
	u.GET("/find/", s.handler.GetUser)
	u.PUT("/delete/", s.handler.DeleteUser)
}
