package chi

type Config struct {
	Host        string `yaml:"host" env:"HOST" env-default:"localhost" env-description:"Host is used to attach server" validate:"hostname"`
	Port        int    `yaml:"port" env:"PORT" env-default:"3000" env-description:"Port is used to attach server" validate:"gte=1,lte=65535"`
	BasePath    string `yaml:"base_path" env:"BASE_PATH" env-default:"/" env-description:"BasePath is used to set base path for all routes"`
	DocsEnabled bool   `yaml:"docs_enabled" env:"DOCS_ENABLED" env-default:"false" env-description:"DocsEnabled is used to enable or disable docs route"`
}
