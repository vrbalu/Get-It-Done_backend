package config

import (
	"log"
	"os"
)

var (
	DBUrl  string
	DBName string
)

func init() {
	loadEnvValues()
}
func loadEnvValues() {
	DBName, DBUrl = loadDBCredentials("MONGODB_NAME", "MONGODB_URL")
}

func loadDBCredentials(dbNameKey string, dbUrlKey string) (string, string) {
	name := os.Getenv(dbNameKey)
	url := os.Getenv(dbUrlKey)
	if name == "" || url == "" {
		log.Fatalf("Env vars %s or %s not set", dbNameKey, dbUrlKey)
	}
	return name, url
}
