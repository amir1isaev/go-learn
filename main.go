package main

import (
	"context"
	"flag"
	"fmt"
	"go-learn/internal/config"
	"go-learn/internal/core"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

var (
	envFile      string
	configFolder string
	configFile   string
)

func init() {
	flag.StringVar(&configFolder, "folder", "configs", "")
	flag.StringVar(&configFile, "file", "main", "")
	flag.StringVar(&envFile, "env", "", "")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer cancel()

	// Initialize env file if exists
	if err := loadEnvFile(); err != nil {
		fmt.Println("[ENV] | [Load]: error occured: %s", err.Error())
	}

	// Initialize config
	if err := config.Init(configFolder, configFile); err != nil {
		fmt.Println("[Config] | [Init]: error occured: %s", err.Error())
	}

	app := core.New(ctx)

	if err := app.Run(); err != nil {
		fmt.Println("[Core] | [Run]: error occured: %s", err.Error())
	}

	if err := app.Stop(); err != nil {
		fmt.Println("[Core] | [Stop]: error occured: %s", err.Error())
	}

}

func loadEnvFile() error {
	if envFile != "" {
		return godotenv.Load(envFile)
	}

	return nil
}
