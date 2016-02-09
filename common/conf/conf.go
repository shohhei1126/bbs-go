package conf

import "github.com/kelseyhightower/envconfig"

type Conf struct {
	DbMaster string `envconfig:"db_master"`
	DbSlave  string `envconfig:"db_slave"`
	LogFile  string `envconfig:"log_file"`
	LogLevel string `envconfig:"log_level"`
	Assets   string `envconfig:"assets"`
}

func Parse() (*Conf, error) {
	appConf := &Conf{}
	err := envconfig.Process("bbsgo", appConf)
	if err != nil {
		panic(err)
	}
	return appConf, nil
}
