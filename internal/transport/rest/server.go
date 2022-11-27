package rest

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/usersService/internal/service"
	"github.com/vslitvinov/usersService/internal/transport/rest/handler"
)



type Server struct {
	Auth *handler.AuthHandler
	Account *handler.AccountHandler
	Server *http.Server
}


func NewServer(db *pgxpool.Pool) *Server {

	serviceAuth := service.NewAuthService(db)
	serviceAccount := service.NewAccountService(db)

	r := &http.ServeMux{}


	s := &Server{
		Auth: handler.NewAuthHandler(serviceAuth),
		Account: handler.NewAccountHandler(serviceAccount),
		Server: &http.Server{},
	}

	s.Server.Handler = r


	r.HandleFunc("/api/v1/auth/signin",s.Auth.SingIn)
	r.HandleFunc("/api/v1/auth/signup",s.Auth.SingUp)
	r.HandleFunc("/api/v1/auth/logout",s.Auth.LogOut)

	return s

}

func (s *Server) Start(addr string)  {
	s.Server.Addr = addr
	go func (s *http.Server){
		s.ListenAndServe()
	}(s.Server)
}
