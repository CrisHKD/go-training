package configg7

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Port  int
	DBUrl string
	Env   string
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig() *Config {
	once.Do(loadConfig)
	return instance
}

func loadConfig() {
	fmt.Println("Cargando datos...")
	time.Sleep(300 * time.Millisecond)

	instance = &Config{
		Port:  8000,
		DBUrl: "postgres://user:pass@localhost:5432/app",
		Env:   "dev",
	}
}