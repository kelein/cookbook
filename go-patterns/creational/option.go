package creational

// Config stands for a server config
type Config struct {
	Timeout  int
	MaxConns int
	Debug    bool
}

// NewConfig create a config instance
func NewConfig(opts ...Option) *Config {
	conf := &Config{
		MaxConns: 1000,
		Timeout:  30,
		Debug:    false,
	}
	for _, opt := range opts {
		opt.Apply(conf)
	}
	return conf
}

// Option setup field for Config
type Option interface {
	Apply(*Config)
}

type optionFunc func(*Config)

func (f optionFunc) Apply(c *Config) { f(c) }

// WithTimeout setup timeout
func WithTimeout(timeout int) Option {
	return optionFunc(func(c *Config) {
		c.Timeout = timeout
	})
}

// WithDebug setup debug mode
func WithDebug() Option {
	return optionFunc(func(c *Config) {
		c.Debug = true
	})
}

// WithMaxConns setup max connections
func WithMaxConns(maxConns int) Option {
	return optionFunc(func(c *Config) {
		c.MaxConns = maxConns
	})
}
