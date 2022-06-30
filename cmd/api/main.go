package main

import (
	"context"
	"github.com/yerlanov/xmexercise/internal"
	"github.com/yerlanov/xmexercise/internal/company/storage"
	"github.com/yerlanov/xmexercise/internal/config"
	"github.com/yerlanov/xmexercise/pkg/postgresql"
	"log"
)

func main() {
	log.Println("config initializing")
	cfg := config.GetConfig("config.yml")

	pg, err := postgresql.NewClient(context.TODO(), 3, cfg.DB.URI)
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}

	store := storage.NewStore(pg)

	server := internal.NewServer(cfg, store)

	if err := server.Start(cfg.Listen.Port); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
