package config

import (
	"fmt"
	"sync"

	"github.com/rasatmaja/mura"
)

var instance *ENV
var singleton sync.Once

// Config ...
type Config struct {
	Filename string
	Type     string
	Path     string
}

// ENV is a stuct to hold all environemnt variable for this app
type ENV struct {
	// Server
	ServerHost       string `mapstructure:"SERVER_HOST" default:"localhost"`
	ServerPort       int    `mapstructure:"SERVER_PORT" default:"2929"`
	ServerReadTO     int    `mapstructure:"SERVER_READ_TIMEOUT" default:"10"`
	ServerWriteTO    int    `mapstructure:"SERVER_WRITE_TIMEOUT" default:"10"`
	ServerIdleTO     int    `mapstructure:"SERVER_IDLE_TIMEOUT" default:"10"`
	ServerProduction bool   `mapstructure:"SERVER_PRODUCTION" default:"false"`
}

// LoadENV ...
func LoadENV() *ENV {
	singleton.Do(func() {
		fmt.Println("[ CNFG ] Starting ENV config ...")
		config := &Config{
			Filename: "app",
			Type:     "env",
			Path:     ".",
		}
		instance = config.BuildENV()
	})
	return instance
}

// BuildENV ...
func (cfg *Config) BuildENV() *ENV {
	env := &ENV{}

	vpr := mura.New()
	vpr.AddConfigPath(cfg.Path)
	vpr.SetConfigName(cfg.Filename)
	vpr.SetConfigType(cfg.Type)

	vpr.FillDefault(env)
	vpr.AutomaticEnv()
	vpr.BindSysEnv(env)
	vpr.ReadInConfig()

	if err := vpr.Unmarshal(&env); err != nil {
		panic(err)
	}

	return env
}
