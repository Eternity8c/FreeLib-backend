package config

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBname      string
	SSLMode     string
	MaxAttempts int
}

func LoadConfig() Config {
	return Config{
		Host:        "localhost",
		Port:        "5432",
		Username:    "admin",
		Password:    "1234",
		DBname:      "my_db",
		SSLMode:     "disable",
		MaxAttempts: 5,
	}
}

func LoadConfigTest() Config {
	return Config{
		Host:        "localhost",
		Port:        "5432",
		Username:    "postgres",
		Password:    "test",
		DBname:      "testdb",
		SSLMode:     "disable",
		MaxAttempts: 5,
	}
}