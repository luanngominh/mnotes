package main

import (
	"fmt"

	"github.com/luanngominh/mnotes/backend/config"
	_ "github.com/luanngominh/mnotes/backend/config"
)

func main() {
	// _, deferDBFunc := db.New(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	// 	config.Cfg.DBHost, config.Cfg.DBPort,
	// 	config.Cfg.DBUser, config.Cfg.DBName,
	// 	config.Cfg.DBPassword, "disable"))

	// defer deferDBFunc()
	fmt.Printf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Cfg.DBHost, config.Cfg.DBPort,
		config.Cfg.DBUser, config.Cfg.DBName,
		config.Cfg.DBPassword, "disable")
}
