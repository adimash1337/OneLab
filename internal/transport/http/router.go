package http

import (
	_ "awesomeProject/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) UserRoutes(incomingRoutes *gin.Engine) {
	//users
	u := incomingRoutes.Group("/user")
	u.POST("/create", s.handler.CreateUser())            //+
	u.POST("/login", s.handler.Login())                  //+
	u.GET("/delete", s.handler.DeleteUser())             //+
	u.PUT("/update_pass", s.handler.UpdatePassword())    //+
	u.GET("/get_by_name", s.handler.GetUserByUsername()) //+
	u.GET("/get_by_id", s.handler.GetUserByID())         //+

	//books
	b := incomingRoutes.Group("/book")
	b.POST("create", s.handler.CreateBook())
	b.GET("/delete", s.handler.DeleteBook())
	b.GET("/get_by_author", s.handler.GetBookByAuthor())
	b.GET("/get_by_id", s.handler.GetBookByID())
	b.GET("/list", s.handler.ListBooks())

	//notes
	n := incomingRoutes.Group("/notes")
	n.POST("/create/:uid/:bid", s.handler.CreateNote())
	n.GET("/delete/:id", s.handler.DeleteNote())
	n.GET("/get_by_id/:id", s.handler.GetByID())
	n.GET("/list_notes", s.handler.ListNotes())

	//swagger
	incomingRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
