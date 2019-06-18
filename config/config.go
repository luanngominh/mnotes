package config

import (
	"crypto/rsa"
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func init() {
	Cfg.DBHost = os.Getenv("DB_HOST")
	Cfg.DBName = os.Getenv("DB_NAME")
	Cfg.DBPort = os.Getenv("DB_PORT")
	Cfg.DBUser = os.Getenv("DB_USER")
	Cfg.DBPassword = os.Getenv("DB_PASSWORD")

	Cfg.Port = os.Getenv("PORT")
	Cfg.Address = os.Getenv("ADDRESS")

	//Load rsa file to signed and verify
	signBytes, err := ioutil.ReadFile(os.Getenv("PRIVATE_KEY_FILE"))
	if err != nil {
		panic(err)
	}

	Cfg.SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_KEY_FILE"))
	if err != nil {
		panic(err)
	}

	Cfg.VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

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
	//SignKey is private key
	//VerifyKey is public key
	//We use RSA for signed
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey

	//Sendgird API key
	SendgridAPI string
}

var (
	//Cfg is global config
	Cfg = &Config{}
)
