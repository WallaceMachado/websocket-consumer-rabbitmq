package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var (
	Port = ""

	RabbitmqUser           = ""
	RabbitmqPass           = ""
	RabbitmqVhost          = ""
	RabbitmqHost           = ""
	RabbitmqPort           = ""
	RabbitmqExchangePerson = ""
	RabbitmqQueuePerson    = ""
)

func Init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := godotenv.Load(basepath + "/../../.env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	Port = os.Getenv("PORT")

	RabbitmqUser = os.Getenv("RABBITMQ_DEFAULT_USER")
	RabbitmqPass = os.Getenv("RABBITMQ_DEFAULT_PASS")
	RabbitmqVhost = os.Getenv("RABBITMQ_DEFAULT_VHOST")
	RabbitmqHost = os.Getenv("RABBITMQ_DEFAULT_HOST")
	RabbitmqPort = os.Getenv("RABBITMQ_DEFAULT_PORT")
	RabbitmqExchangePerson = os.Getenv("RABBITMQ_EXCHANGE_PERSON")
	RabbitmqQueuePerson = os.Getenv("RABBITMQ_QUEUE_PERSON")

}
