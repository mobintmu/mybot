package app

// Application interface for start application
type Application interface {
	GoDotEnvVariable(key string) string
}

type App struct{}
