package config

type Config struct {
	PostgresConnStr string
	GrpcPort        int
	RestPort        int
}

func Default() *Config {
	return &Config{
		PostgresConnStr: "user=postgres password=postgres dbname=postgres sslmode=disable",
		GrpcPort:        9090,
		RestPort:        8000,
	}
}
