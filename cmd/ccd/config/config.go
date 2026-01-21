package config

import (
	"github.com/indexdata/ccms/internal/global"
	"gopkg.in/ini.v1"
)

type Config struct {
	DB *Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func New(datadir string) (*Config, error) {
	c, err := ini.Load(global.ServerConfigFileName(datadir))
	if err != nil {
		return nil, err
	}
	s := c.Section("main")

	return &Config{
		DB: &Database{
			Host:     s.Key("host").String(),
			Port:     s.Key("port").String(),
			User:     s.Key("user").String(),
			Password: s.Key("password").String(),
			DBName:   s.Key("dbname").String(),
			SSLMode:  s.Key("sslmode").String(),
		},
	}, nil
}

func InitStub() string {
	return "[main]\n" +
		"host = \n" +
		"port = 5432\n" +
		"user = ccms\n" +
		"password = \n" +
		"dbname = \n" +
		"sslmode = require\n"
}

func (d *Database) ConnString() string {
	return "connect_timeout=30 host=" + d.Host + " port=" + d.Port +
		" user=" + d.User + " password=" + d.Password +
		" dbname=" + d.DBName + " sslmode=" + d.SSLMode
}
