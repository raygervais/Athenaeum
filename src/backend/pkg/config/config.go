package config

type Config struct {
	dbLive     string
	dbTest     string
	dbDev      string
	serverPort string
}

func Get() *Config {
	conf := &Config{
		dbDev:      "dev.db",
		dbLive:     "live.db",
		dbTest:     "test.db",
		serverPort: "8000",
	}

	return conf
}

func (c *Config) GetDatabaseName(environment string) string {
	switch environment {
	case "live":
		return c.dbLive
	case "test":
		return c.dbTest
	default:
		return c.dbTest
	}
}

func (c *Config) GetAPIPort() string {
	return ":" + c.serverPort
}
