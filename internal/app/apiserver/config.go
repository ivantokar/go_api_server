package apiserver

import "github.com/ivantokar/GO_PRACTICE/internal/app/store"

type Config struct {
    BindAddr string `toml:"bing_addr"`
    LogLevel string `toml:"log_level"`
    Store    *store.Config
}

func NewConfig() *Config {
    return &Config {
        BindAddr: ":8080",
        LogLevel: "debug",
        Store:    store.NewConfig(),
    }
}
