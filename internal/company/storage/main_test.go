package storage

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/yerlanov/xmexercise/internal/config"
	"github.com/yerlanov/xmexercise/pkg/postgresql"
	"log"
	"os"
	"testing"
)

var queries *Queries
var pg *pgxpool.Pool

func TestMain(m *testing.M) {
	cfg := config.GetConfig("../../../config.yml")

	var err error
	pg, err = postgresql.NewClient(context.TODO(), 3, cfg.DB.URI)
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}

	queries = New(pg)
	os.Exit(m.Run())
}
