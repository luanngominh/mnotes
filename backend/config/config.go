package config

import "os"

func init() {
	Cfg.DBHost = os.Getenv("DB_HOST")
	Cfg.DBName = os.Getenv("DB_NAME")
	Cfg.DBPort = os.Getenv("DB_PORT")
	Cfg.DBUser = os.Getenv("DB_USER")
	Cfg.DBPassword = os.Getenv("DB_PASSWORD")

	Cfg.Port = os.Getenv("PORT")
	Cfg.Address = os.Getenv("ADDRESS")

	Cfg.PrivateKeyFile = os.Getenv("PRIVATE_KEY_FILE")
	Cfg.PublicKeyFile = os.Getenv("PUBLIC_KEY_FILE")

	Cfg.SendgridAPI = os.Getenv("SENDGIRD_APIKEY")
}

//Config store program config
type Config struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBUser     string
	DBPassword string

	//Port running port
	Port string
	//Server will listen on ip adress
	Address string

	//Use for jwt
	PrivateKeyFile string
	PublicKeyFile  string

	//Sendgird API key
	SendgridAPI string
}

var (
	//Cfg is global config
	Cfg = &Config{}
)
