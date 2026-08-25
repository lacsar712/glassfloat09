package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"example.com/glassfloat/internal/config"
	"example.com/glassfloat/internal/db"
	"example.com/glassfloat/internal/policy"
	"example.com/glassfloat/internal/store"
	"example.com/glassfloat/internal/svc"
)

func main() {
	cfg := config.Config{DBPath: "glassfloat.sqlite", DataDir: "data", AdminToken: "dev-token"}.Normalized()
	_ = os.MkdirAll(cfg.DataDir, 0o755)
	sqlDB, err := db.Open(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	cat := svc.NewCatalog(store.New(sqlDB), cfg)
	ctx := context.Background()
	if err := cat.BootstrapAdmin(ctx); err != nil {
		log.Fatal(err)
	}
	items := policy.Seed()
	for _, it := range items {
		if _, err := cat.Create(ctx, it.Title, it.Body, it.Tags, 1); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("seeded", len(items))
}
