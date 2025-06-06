package config

type Config struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() Config {
	return Config{
		Email:    "you@example.com",
		Password: "1234",
		Address:  "my@example.com",
	}
}
