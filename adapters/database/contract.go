package database

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

type Provider[K any] interface {
	Exec() K
	Connect(cfg Config) error
	Close() error
}
