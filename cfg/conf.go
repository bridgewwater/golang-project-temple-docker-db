package cfg

type ConfFile struct {
	ConfigPath   string
	WriteLogFile bool
}

var global *Conf

func Global() Conf {
	return *global
}

type Conf struct {
	RunMode      string        `json:"run_mode" mapstructure:"run_mode"`
	Name         string        `json:"name" mapstructure:"name"`
	Addr         string        `json:"addr" mapstructure:"addr"`
	SslEnable    bool          `json:"ssl_enable" mapstructure:"ssl_enable"`
	Zap          Zap           `json:"zap" mapstructure:"zap"`
	RedisClients []RedisClient `json:"redis_clients" mapstructure:"redis_clients"`
	MysqlClients []MysqlClient `json:"mysql_clients" mapstructure:"mysql_clients"`
}