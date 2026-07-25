package configs

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:",squash"`
	Redis    RedisConfig    `mapstructure:",squash"`
	RabbitMQ RabbitMQConfig `mapstructure:",squash"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

type RedisConfig struct {
	URL string
}

type RabbitMQConfig struct {
}
