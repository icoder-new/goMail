package utils

import (
	"fr33d0mz/goMail/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var AppSettings models.Settings

func PutAdditionalSettings() {
	AppSettings.AppParams.LogDebug = "./logs/debug.log"
	AppSettings.AppParams.LogInfo = "./logs/info.log"
	AppSettings.AppParams.LogWarning = "./logs/warning.log"
	AppSettings.AppParams.LogError = "./logs/error.log"

	AppSettings.AppParams.LogCompress = true
	AppSettings.AppParams.LogMaxSize = 10
	AppSettings.AppParams.LogMaxAge = 100
	AppSettings.AppParams.LogMaxBackups = 100
}

func ReadSettings() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read `config.yaml` file. Error is: ", err.Error())
	}

	AppSettings.AppParams.PortRun = viper.GetString("port")

	AppSettings.PostgresParams.User = viper.GetString("db.username")
	AppSettings.PostgresParams.Password = os.Getenv("DB_PASSWORD")
	AppSettings.PostgresParams.Server = viper.GetString("db.host")
	AppSettings.PostgresParams.Port = viper.GetString("db.port")
	AppSettings.PostgresParams.Database = viper.GetString("db.dbname")
	AppSettings.PostgresParams.SSLMode = viper.GetString("db.sslmode")

	AppSettings.SMTPParams.Server = viper.GetString("smtp.host")
	AppSettings.SMTPParams.Port = viper.GetString("smtp.port")
	AppSettings.SMTPParams.Username = viper.GetString("smtp.username")
	AppSettings.SMTPParams.Password = os.Getenv("SMTP_PASSWORD")
}
