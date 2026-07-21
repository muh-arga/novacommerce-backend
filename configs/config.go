package configs

type Config struct {
	App      AppConfig      `mapstructure:",squash"`
	Database DatabaseConfig `mapstructure:",squash"`
	Redis    RedisConfig    `mapstructure:",squash"`
	RabbitMQ RabbitMQConfig `mapstructure:",squash"`
}

type AppConfig struct {
	Name string `mapstructure:"APP_NAME"`
	Port int    `mapstructure:"APP_PORT"`
	Env  string `mapstructure:"APP_ENV"`
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
