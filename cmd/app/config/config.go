package config

type Config struct {
	Default DefaultOptions `yaml:"default"`
	Mysql   MysqlOptions   `yaml:"mysql"`
}

type DefaultOptions struct {
	Listen int    `yaml:"listen"`
	LogDir string `yaml:"log_dir"`
}

type MysqlOptions struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}
