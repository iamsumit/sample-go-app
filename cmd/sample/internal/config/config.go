package config

type Configuration struct {
	LaunchDarkly LaunchDarklyConfig `mapstructure:"launch_darkly"`
	MySQL        MySQLConfig
	Environment  EnvironmentConfig
	Jaeger       JaegerConfig
}

type LaunchDarklyConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

type MySQLConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type EnvironmentConfig struct {
	Env      string
	OtherEnv string
	EmptyEnv string
}

type JaegerConfig struct {
	Host string
	Path string
}
