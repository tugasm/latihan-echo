package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
	connStr := "postgresql://postgres.kltjxgpctqsrftstvqtn:abc1029384756@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parsing config DB: %v", err)
	}

	config.ConnConfig.ConnectTimeout = 5 * time.Second

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Failed to create db pooling: %v", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("DB failed Ping: %v", err)
	}

	fmt.Println("Database connected")
}

func CloseDB() {
	Pool.Close()
}
