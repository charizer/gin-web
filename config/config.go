package config

import (
	"github.com/spf13/viper"
	"strings"
)

func init() {
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetDefault("HTTP_PORT", 12035)
	viper.SetDefault("MODE", "debug")
	viper.SetDefault("MONGO_URL", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_DB", "gin-web")
	viper.SetDefault("DEBUG_LOG", true)
	viper.SetDefault("TRACE_LOG", true)
	viper.SetDefault("LOG_FILE", "./logs/gin-web.log")
	viper.SetDefault("REDIS_URL", "['redis://localhost:6379/0']")
}

type Config struct {
	Port     string
	Mode     string
	MongoURL string
	MongoDB  string
	DebugLog bool
	TraceLog bool
	LogFile  string
	RedisURL []string
}

func GetConfig() *Config {
	return &Config{
		Port:     viper.GetString("HTTP_PORT"),
		Mode:     viper.GetString("MODE"),
		MongoURL: viper.GetString("MONGO_URL"),
		MongoDB:  viper.GetString("MONGO_DB"),
		DebugLog: viper.GetBool("DEBUG_LOG"),
		TraceLog: viper.GetBool("TRACE_LOG"),
		LogFile:  viper.GetString("LOG_FILE"),
		RedisURL: viper.GetStringSlice("REDIS_URL"),
	}
}
