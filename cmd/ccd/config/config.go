package config

import (
	"fmt"
	"strings"

	"github.com/indexdata/ccms/internal/crypto"
	"github.com/indexdata/ccms/internal/global"
	"gopkg.in/ini.v1"
)

type Config struct {
	DB       *Database
	Security *Security
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Security struct {
	SecretKey []byte
}

func New(datadir string) (*Config, error) {
	c, err := ini.Load(global.ServerConfigFileName(datadir))
	if err != nil {
		return nil, err
	}

	s := c.Section("security")
	secretKey, err := readSecretKey(s.Key("secret_key").String())
	if err != nil {
		return nil, err
	}
	security := &Security{
		SecretKey: secretKey,
	}

	s = c.Section("db.main")
	db := &Database{
		Host:     s.Key("host").String(),
		Port:     s.Key("port").String(),
		User:     s.Key("user").String(),
		Password: s.Key("password").String(),
		DBName:   s.Key("dbname").String(),
		SSLMode:  s.Key("sslmode").String(),
	}

	return &Config{
		DB:       db,
		Security: security,
	}, nil
}

func InitStub() string {
	key := crypto.EncodeToHexString(crypto.RandomKey())
	return "[db.main]\n" +
		"host = \n" +
		"port = 5432\n" +
		"user = ccms\n" +
		"password = \n" +
		"dbname = \n" +
		"sslmode = require\n" +
		"\n" +
		"[security]\n" +
		"secret_key = " + key + "\n" +
		"\n"
}

func (d *Database) ConnString() string {
	return "connect_timeout=30 host=" + d.Host + " port=" + d.Port +
		" user=" + d.User + " password=" + d.Password +
		" dbname=" + d.DBName + " sslmode=" + d.SSLMode
}

func readSecretKey(key string) ([]byte, error) {
	if strings.TrimSpace(key) == "" {
		return nil, fmt.Errorf("secret key not configured")
	}
	k, err := crypto.DecodeFromHexString(key)
	if err != nil {
		return nil, fmt.Errorf("reading secret key: %v", err)
	}
	return k, nil
}
