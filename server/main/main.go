package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"plan-O/repository/pg"
	"plan-O/server"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	repo, err := pg.New("postgresql://postgres:psql@localhost:5432/crm")
	if err != nil {
		panic(err)
	}
	s := server.New(repo)
	go func() {
		err := s.Run("localhost:8085")
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
