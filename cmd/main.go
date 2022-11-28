package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/vslitvinov/usersService/internal/service"
	"github.com/vslitvinov/usersService/internal/transport/rest/handler"
	"github.com/vslitvinov/usersService/pkg/helpers"
)


func main(){
	dbConfig := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		"postgres",
		"",
		"localhost",
		"5433",
		"ptom",
	)
	
	serverConfig := fmt.Sprintf("%v:%v",
		"localhost",
		"8080",
	)

	dbPool, err  := helpers.NewPostgres(context.TODO(),dbConfig)
	if err != nil {
		log.Println(err)
	}


	accountService := service.NewAccountService(dbPool)
	sessionService := service.NewSessionStorage(dbPool)
	authService := service.NewAuthService(accountService,sessionService)


	router := http.NewServeMux()

	authHandler := handler.NewAuthHandler(authService)

	router.HandleFunc("/singin",authHandler.SingIn)
	router.HandleFunc("/singup",authHandler.SingUp)

	server := http.Server{
		Addr: serverConfig,
		Handler: router,
	}

	server.ListenAndServe()


}