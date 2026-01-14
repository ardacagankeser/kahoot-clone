package config

// Config holds application configuration loaded from environment variables.
// This is a placeholder for future environment loading functionality.
type Config struct {
	MongoURI   string
	ServerPort string
}

// Load returns a Config with default values.
// Future: load from environment variables.
func Load() *Config {
	return &Config{
		MongoURI:   "mongodb://localhost:27017",
		ServerPort: ":3000",
	}
}
