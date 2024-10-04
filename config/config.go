package config
 
import (
    "log"
    "github.com/spf13/viper"
)
 
type Config struct {
    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Name     string `mapstructure:"name"`
    } `mapstructure:"database"`
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`
}
 
var AppConfig Config
 
// LoadConfig reads the configuration file and unmarshals it into AppConfig
func LoadConfig() {
    viper.SetConfigName("config") // name of config file (without extension)
    viper.SetConfigType("toml")   // required if the config file does not have the extension
    viper.AddConfigPath(".")      // optionally look for config in the working directory
 
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %s", err)
    }
 
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode into struct: %v", err)
    }
}
 
 