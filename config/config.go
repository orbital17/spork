package config

type Config struct {
	PostgresConnStr string
	GrpcPort        int
}

func Default() *Config {
	return &Config{
		PostgresConnStr: "user=postgres password=postgres dbname=postgres sslmode=disable",
		GrpcPort:        8080,
	}
}
