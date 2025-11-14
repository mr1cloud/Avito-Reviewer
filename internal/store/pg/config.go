package pg

// Config is struct to configure service postgres store
type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost" env-description:"Postgres database host"`
	Port     int    `yaml:"port" env:"PORT" env-default:"5432" env-description:"Postgres database port" validate:"gte=1,lte=65535"`
	User     string `yaml:"user" env:"USER" env-default:"postgres" env-required:"true" env-description:"Postgres database user"`
	Password string `yaml:"password" env:"PASSWORD" env-required:"true" env-description:"Postgres database password for user"`
	DB       string `yaml:"db" env:"DATABASE" env-default:"postgres" env-description:"Postgres database name"`
}
