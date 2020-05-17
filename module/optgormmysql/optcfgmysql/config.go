package optcfgmysql

// mysql config for mysql
type Config struct {
	Name            string `json:"name" toml:"name" mapstructure:"name" yaml:"name"`
	OpenLog         string `json:"open_log" toml:"open_log" mapstructure:"open_log" yaml:"open_log"`
	Host            string `json:"host" toml:"host" mapstructure:"host" yaml:"host"`
	DBName          string `json:"db_name" toml:"db_name" mapstructure:"db_name" yaml:"db_name"`
	Username        string `json:"username" toml:"username" mapstructure:"username" yaml:"username"`
	Password        string `json:"password" toml:"password" mapstructure:"password" yaml:"password"`
	Config          string `json:"config" toml:"config" mapstructure:"config" yaml:"config"`
	MaxIdleConns    int    `json:"max_idle_conns" toml:"max_idle_conns" mapstructure:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns" toml:"max_open_conns" mapstructure:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" toml:"conn_max_lifetime" mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime"`
}

func ByName(list []Config, name string) Config {
	for _, client := range list {
		if client.Name == name {
			return client
		}
	}
	return Config{}
}
