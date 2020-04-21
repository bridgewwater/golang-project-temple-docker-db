package cfg

type Zap struct {
	AtomicLevel   int           `json:"AtomicLevel" mapstructure:"AtomicLevel"`
	Development   bool          `json:"Development" mapstructure:"Development"`
	EncoderConfig EncoderConfig `json:"EncoderConfig" mapstructure:"EncoderConfig"`
	Encoding      string        `json:"Encoding" mapstructure:"Encoding"`
	Fields        Fields        `json:"Fields" mapstructure:"Fields"`
	FieldsAuto    bool          `json:"FieldsAuto" mapstructure:"FieldsAuto"`
	Rotate        Rotate        `json:"rotate"  mapstructure:"rotate"`
}

type EncoderConfig struct {
	CallerKey      string `json:"CallerKey" mapstructure:"CallerKey"`
	LevelKey       string `json:"LevelKey" mapstructure:"LevelKey"`
	MessageKey     string `json:"MessageKey" mapstructure:"MessageKey"`
	NameKey        string `json:"NameKey" mapstructure:"NameKey"`
	StacktraceKey  string `json:"StacktraceKey"  mapstructure:"StacktraceKey"`
	TimeKey        string `json:"TimeKey" mapstructure:"TimeKey"`
	TimeEncoder    string `json:"TimeEncoder" mapstructure:"TimeEncoder"`
	EncodeDuration string `json:"EncodeDuration" mapstructure:"EncodeDuration"`
	EncodeLevel    string `json:"EncodeLevel" mapstructure:"EncodeLevel"`
	EncodeCaller   string `json:"EncodeCaller" mapstructure:"EncodeCaller"`
}

type Fields struct {
	Key string `json:"Key" mapstructure:"Key"`
	Val string `json:"Val" mapstructure:"Val"`
}

type Rotate struct {
	Compress   bool   `json:"Compress" mapstructure:"Compress"`
	Filename   string `json:"Filename" mapstructure:"Filename"`
	MaxAge     int    `json:"MaxAge" mapstructure:"MaxBackups"`
	MaxBackups int    `json:"MaxBackups" mapstructure:"MaxAge"`
	MaxSize    int    `json:"MaxSize" mapstructure:"Compress"`
}
