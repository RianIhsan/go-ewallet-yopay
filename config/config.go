package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	AppPort  int
	Secret   string
	Database database
	Midtrans midtrans
}

type database struct {
	DbHost string
	DbPort int
	DbUser string
	DbPass string
	DbName string
}

type midtrans struct {
	ClientKey string
	ServerKey string
}

func loadConfig() *Config {
	var res = new(Config)
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Failed to fetch .env file")
		}
	}

	if value, found := os.LookupEnv("PORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid server port", err.Error())
			return nil
		}
		res.AppPort = port
	}

	if value, found := os.LookupEnv("SECRET"); found {
		res.Secret = value
	}

	if value, found := os.LookupEnv("DBHOST"); found {
		res.Database.DbHost = value
	}

	if value, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Config : invalid db port", err.Error())
			return nil
		}
		res.Database.DbPort = port
	}

	if value, found := os.LookupEnv("DBUSER"); found {
		res.Database.DbUser = value
	}

	if value, found := os.LookupEnv("DBPASS"); found {
		res.Database.DbPass = value
	}

	if value, found := os.LookupEnv("DBNAME"); found {
		res.Database.DbName = value
	}

	if value, found := os.LookupEnv("MCKEY"); found {
		res.Midtrans.ClientKey = value
	}

	if value, found := os.LookupEnv("MSKEY"); found {
		res.Midtrans.ServerKey = value
	}

	return res
}

func BootConfig() *Config {
	return loadConfig()
}
