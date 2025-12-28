package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pnkrshv/REST-API-GO/internal/config"
)

func main() {
	cfg := config.MustLoad()

	//TODO: logger: slog

	//TODO: storage: sqlite (postgre)

	//TODO: router: chi (echo)

	//TODO: run server
}

func setupLogger(env string) {
	var log *slog.Logger

	switch env {
	case "local": 
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	}
}
