package config

type Config struct {
	Db ConnectorIface
}

func InitConfig() *Config {
	init := new(Config)
	init.NewDb()
	return init
}

func (c *Config) NewDb() {
	c.Db = InitDB()
}
