package main

import (
	"fmt"

	"github.com/kekaswork/TradeSim/backend/services/auth/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	_ = cfg

	fmt.Printf("config: %v", cfg)
}
