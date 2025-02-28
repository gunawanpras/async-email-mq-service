package config

type (
	Config struct {
		Server   ServerConfig `yaml:"server"`
		Postgre  PostgreList  `yaml:"postgre"`
		RabbitMQ RabbitMQList `yaml:"rabbitmq"`
	}

	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	PostgreList struct {
		Primary DatabaseConfig `yaml:"primary"`
	}

	DatabaseConfig struct {
		ConnString              string `yaml:"connString"`
		MaxOpenConn             int    `yaml:"maxOpenConn"`
		MaxIdleConn             int    `yaml:"maxIdleConn"`
		MaxConnLifeTimeInSecond int    `yaml:"maxConnLifeTimeInSecond"`
	}

	RabbitMQList struct {
		Primary RabbitMQConfig `yaml:"primary"`
	}

	RabbitMQConfig struct {
		ConnString string `yaml:"connString"`
	}
)
