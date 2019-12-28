package apiserver

type Config struct {
    BindAddr string `toml:"bing_addr"`
    LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
    return &Config {
        BindAddr: ":8080",
        LogLevel: "debug",
    }
}
