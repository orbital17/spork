package config

type Config struct {
	PostgresConnStr string
}

func Default() *Config {
	return &Config{
		PostgresConnStr: "user=postgres password=postgres dbname=postgres sslmode=disable",
	}
}
