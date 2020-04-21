package cfg

type MysqlClient struct { // mysql
	Name            string `json:"name" mapstructure:"name"`
	OpenLog         bool   `json:"open_log" mapstructure:"open_log"`
	Username        string `json:"username" mapstructure:"username"`
	Password        string `json:"password" mapstructure:"password"`
	Host            string `json:"host" mapstructure:"host"`
	Dbname          string `json:"dbname" mapstructure:"dbname"`
	Config          string `json:"config" mapstructure:"config"`
	MaxIdleConns    int    `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `json:"max_open_conns" mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
}
