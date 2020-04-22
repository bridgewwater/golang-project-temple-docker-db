package cfg

type RedisClient struct {
	Name         string `json:"name" mapstructure:"name"`
	Addr         string `json:"addr" mapstructure:"addr"`
	Password     string `json:"password" mapstructure:"password"`
	DB           int    `json:"db" mapstructure:"db"`
	MaxRetries   int    `json:"max_retries" mapstructure:"max_retries"`
	DialTimeout  int    `json:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  int    `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout int    `json:"write_timeout" mapstructure:"write_timeout"`
}

func RedisClientByName(name string) RedisClient {
	conf := *global
	for _, client := range conf.RedisOptConfig {
		if client.Name == name {
			return RedisClient{}
		}
	}
	return RedisClient{}
}
