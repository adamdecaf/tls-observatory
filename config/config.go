package config

import (
	"fmt"

	"github.com/jvehent/gozdef"
	"gopkg.in/gcfg.v1"
)

type ObserverConfig struct {
	General struct {
		RabbitMQRelay  string
		Postgres       string
		PostgresPass   string
		PostgresDB     string
		PostgresUser   string
		CipherscanPath string
		GoRoutines     int // * cores = The Max number of spawned Goroutines
	}
	TrustStores struct {
		Name []string
		Path []string
	}
	MozDef gozdef.MqConf
}

type APIConfig struct {
	Postgres     string
	PostgresPass string
}

func APIConfigLoad(path string) (conf APIConfig, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("configLoad() -> %v", e)
		}
	}()
	var c APIConfig
	err = gcfg.ReadFileInto(&c, path)

	return c, err
}

func ObserverConfigLoad(path string) (conf ObserverConfig, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("configLoad() -> %v", e)
		}
	}()
	var c ObserverConfig
	err = gcfg.ReadFileInto(&c, path)

	return c, err
}

func GetObserverDefaults() ObserverConfig {
	conf := ObserverConfig{}

	conf.General.RabbitMQRelay = "amqp://guest:guest@localhost:5672/"
	conf.TrustStores.Name = append(conf.TrustStores.Name, "")
	conf.TrustStores.Path = append(conf.TrustStores.Path, "")
	conf.General.Postgres = "127.0.0.1:5432"
	conf.General.PostgresPass = "password"
	conf.General.PostgresDB = "observer"
	conf.General.PostgresUser = "observer"
	conf.General.CipherscanPath = "../../../cipherscan/cipherscan"
	conf.General.GoRoutines = 10

	return conf
}
