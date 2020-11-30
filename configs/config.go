package configs

import (
	"fmt"
	"os"
	"reflect"
)

// DevConfig Development configuration

type config struct {
	DSN   string
	Debug bool
	Port  string `default:"80"`
	Host  string `default:"localhost"`
}

func (c *config) initDefaults() {
	port, _ := reflect.TypeOf(Config).FieldByName("Port")
	host, _ := reflect.TypeOf(Config).FieldByName("Host")

	Config.Port = port.Tag.Get("default")
	Config.Host = host.Tag.Get("default")
}

// Config for project
var Config config

// DevConfig Development configuration
func DevConfig() {
	Config.DSN = "host=localhost user=postgres password=postgres dbname=company_service port=5432 timezone=Asia/Baku"
	Config.Debug = true
}

// ProdConfig Production configuration
func ProdConfig() {
	Config.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=enable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	Config.Debug = false
}

// Init Config
func init() {
	if s := os.Getenv("settings"); s == "prod" {
		ProdConfig()
	} else {
		DevConfig()
		Config.initDefaults()
	}
}
