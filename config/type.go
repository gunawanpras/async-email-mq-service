package config

type (
	Config struct {
		Server   ServerConfig `yaml:"server"`
		Postgres PostgresList `yaml:"postgres"`
	}

	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	PostgresList struct {
		Primary DatabaseConfig `yaml:"primary"`
	}

	DatabaseConfig struct {
		ConnString              string `yaml:"connString"`
		MaxOpenConn             int    `yaml:"maxOpenConn"`
		MaxIdleConn             int    `yaml:"maxIdleConn"`
		MaxConnLifeTimeInSecond int    `yaml:"maxConnLifeTimeInSecond"`
	}
)
