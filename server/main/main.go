package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"plan-O/repository/pg"
	"plan-O/server"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	repo, err := pg.New("postgresql://postgres:admin@91.142.72.135:5432/crm")
	if err != nil {
		panic(err)
	}
	s := server.New(repo)
	go func() {
		err := s.Run("http://91.142.72.135:80")
		log.Println("server is running")
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	<-ctx.Done()
}
