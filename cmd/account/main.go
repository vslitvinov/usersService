package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vslitvinov/usersService/internal/transport/rest"
	"github.com/vslitvinov/usersService/pkg/helpers"
)

func main(){
	
	dbConfig := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		"root",
		"pAEiYSOu",
		"159.223.19.175",
		"5432",
		"web",
	)
	
	serverConfig := fmt.Sprintf("%v:%v",
		"localhost",
		"8080",
	)

	dbPool, err  := helpers.NewPostgres(context.TODO(),dbConfig)
	if err != nil {
		log.Println(err)
	}

	server := rest.NewServer(dbPool)

	server.Start(serverConfig)


	for {}

}