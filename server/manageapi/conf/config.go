package conf

import "manageapi/conf/config"

type Config struct {
	Mysql  config.Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Common config.Common `mapstructure:"common" json:"common" yaml:"common"`
}
