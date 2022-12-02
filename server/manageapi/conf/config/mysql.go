package config

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}
